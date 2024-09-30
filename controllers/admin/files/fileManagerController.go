package admin

import (
	"fmt"

	"log"
	"net/http"
	"os"
	"project/config"
	"project/internal/layouts"
	"project/internal/links"
	"strings"
	"time"

	"github.com/gouniverse/filesystem"
	"github.com/mingrammer/cfmt"

	"github.com/golang-module/carbon/v2"
	"github.com/gouniverse/api"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
	"github.com/gouniverse/responses"
	"github.com/gouniverse/utils"

	"github.com/hyperjiang/php"

	"github.com/samber/lo"
)

const JSON_ACTION_FILE_RENAME = "file_rename"
const JSON_ACTION_FILE_DELETE = "file_delete"
const JSON_ACTION_FILE_UPLOAD = "file_upload"
const JSON_ACTION_DIRECTORY_CREATE = "directory_create"
const JSON_ACTION_DIRECTORY_DELETE = "directory_delete"
const MAX_UPLOAD_SIZE = 50 * 1024 * 1024 // 50MB

func NewFileManagerController() *FileManagerController {
	// //        $this->user = \App\Helpers\AppHelper::getUser('admin');
	// //        if ($this->user == null) {
	// //            die('User authentication needed to use this service');
	// //            exit;
	// //        }

	//         $this->disk = 'media_manager';
	// //        $rootDir = trim(request('root_dir', ''));
	// //        if ($rootDir == '') {
	// //            die('Root directory is required');
	// //        }
	// //        $this->filesRootDir = public_path('media'); //public_path() . DIRECTORY_SEPARATOR . 'files' . DIRECTORY_SEPARATOR;
	// //        $this->filesRootUrl = url('/') . '/media/';
	// //        $rootDir = trim(request('root_dir', ''));
	// //        if ($rootDir == '') {
	// //            die('Root directory is required');
	// //        }
	// //        $rootDir = trim($rootDir, '/');
	// //        $rootDir = trim($rootDir, '.');
	// //        $this->fileManagerRootDir = $this->filesRootDir . $rootDir . DIRECTORY_SEPARATOR;
	// //        $this->fileManagerRootUrl = $this->filesRootUrl . $rootDir . '/';
	// //
	// //        $dirExists = \Storage::disk($this->disk)->exists($this->fileManagerRootDir);
	// //
	// //        if($dirExists==false){
	// //            $result = \Storage::disk($this->disk)->makeDirectory($this->fileManagerRootDir);
	// //        }
	rootDirPath := strings.TrimSpace(config.MediaRoot)
	rootDirPath = strings.Trim(rootDirPath, "/")
	rootDirPath = strings.Trim(rootDirPath, ".")
	rootDirPath = "/" + rootDirPath

	return &FileManagerController{
		rootDirPath: rootDirPath,
	}
}

type FileEntry struct {
	IsDir             bool
	Path              string
	URL               string
	Name              string
	Size              int64
	SizeHuman         string
	LastModified      time.Time
	LastModifiedHuman string
}

type FileManagerController struct {
	// rootDir if not empty will be used as the root/top directory
	rootDirPath string
	funcLayout  func(content string) string
	storage     filesystem.StorageInterface
}

func (controller *FileManagerController) init(r *http.Request) string {
	// cfmt.Infoln(config.MediaEndpoint)

	// controller.storage, err = filesystem.NewStorage(filesystem.Disk{
	// 	DiskName:             "S3",
	// 	Driver:               filesystem.DRIVER_S3,
	// 	Url:                  config.MediaUrl,
	// 	Region:               config.MediaRegion,
	// 	Key:                  config.MediaKey,
	// 	Secret:               config.MediaSecret,
	// 	Bucket:               config.MediaBucket,
	// 	UsePathStyleEndpoint: true,
	// })

	controller.storage = config.SqlFileStorage

	controller.funcLayout = func(content string) string {
		return layouts.NewAdminLayout(r, layouts.Options{
			Title:   "File Manager",
			Content: hb.Raw(content),
		}).ToHTML()
	}

	return ""
}

func (c *FileManagerController) AnyIndex(w http.ResponseWriter, r *http.Request) string {
	c.init(r)

	if lo.Contains([]string{
		JSON_ACTION_DIRECTORY_CREATE,
		JSON_ACTION_DIRECTORY_DELETE,
		JSON_ACTION_FILE_RENAME,
		JSON_ACTION_FILE_DELETE,
		JSON_ACTION_FILE_UPLOAD,
	}, strings.TrimSpace(utils.Req(r, "action", ""))) {
		responses.JSONResponseF(w, r, c.anyIndex)
		return ""
	}

	responses.HTMLResponseF(w, r, c.anyIndex)
	return ""
}

func (c *FileManagerController) anyIndex(w http.ResponseWriter, r *http.Request) string {
	action := strings.TrimSpace(utils.Req(r, "action", ""))
	if action == JSON_ACTION_FILE_RENAME {
		return c.fileRenameAjax(r)
	}

	if action == JSON_ACTION_FILE_DELETE {
		return c.fileDeleteAjax(r)
	}

	if action == JSON_ACTION_DIRECTORY_CREATE {
		return c.directoryCreateAjax(r)
	}

	if action == JSON_ACTION_DIRECTORY_DELETE {
		return c.directoryDeleteAjax(r)
	}

	if action == JSON_ACTION_FILE_UPLOAD {
		return c.fileUploadAjax(r)
	}

	return c.getMediaManager(r)
}

func (c *FileManagerController) fileUploadAjax(r *http.Request) string {
	if r.ContentLength > MAX_UPLOAD_SIZE {
		return api.Error("The uploaded image is too big. Please use an file less than 50MB in size").ToString()
	}

	currentDir := utils.Req(r, "current_dir", "")
	if currentDir == "" {
		return api.Error("current_dir is required").ToString()
	}

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("upload_file")
	if err != nil {
		return api.Error(err.Error()).ToString()
	}
	defer file.Close() // Cleanup

	filePath, errSave := utils.FileSaveToTempDir(fileHeader.Filename, file)
	if errSave != nil {
		log.Println(errSave.Error())
		return api.Error(errSave.Error()).ToString()
	}
	defer os.Remove(filePath) // Cleanup

	remoteFilePath := currentDir + "/" + fileHeader.Filename

	data, err := os.ReadFile(filePath)
	if err != nil {
		return api.Error(err.Error()).ToString()
	}

	if c.storage == nil {
		return api.Error("Storage not initialized").ToString()
	}

	err = c.storage.Put(remoteFilePath, data)

	if err != nil {
		return api.Error(err.Error()).ToString()
	}

	return api.Success("File uploaded successfully").ToString()
}

func (c *FileManagerController) directoryCreateAjax(r *http.Request) string {
	newDirName := strings.TrimSpace(utils.Req(r, "create_dir", ""))

	if newDirName == "" {
		return api.Error("create_dir is required").ToString()
	}

	currentDir := strings.TrimSpace(utils.Req(r, "current_dir", ""))

	if currentDir == "" {
		return api.Error("current_dir is required").ToString()
	}

	if currentDir == "/" {
		currentDir = "" // to prevent double slashes
	}

	dirPath := currentDir + "/" + newDirName
	dirPath = strings.ReplaceAll(dirPath, "//", "/") // remove double slashes
	dirPath = strings.TrimRight(dirPath, "/")        // remove trailing slashes

	// cfmt.Infoln("New directory:", dirPath)

	if dirPath == "" || dirPath == "/" {
		return api.Error("root directory can not be created").ToString()
	}

	if c.storage == nil {
		return api.Error("Storage not initialized").ToString()
	}

	errDeleted := c.storage.MakeDirectory(dirPath)

	if errDeleted == nil {
		return api.Success("directory created successfully").ToString()
	}

	return api.Error(errDeleted.Error()).ToString()
}

func (c *FileManagerController) directoryDeleteAjax(r *http.Request) string {
	selectedDirName := strings.TrimSpace(utils.Req(r, "delete_dir", ""))

	if selectedDirName == "" {
		return api.Error("delete_dir is required").ToString()
	}

	currentDir := strings.TrimSpace(utils.Req(r, "current_dir", ""))

	if currentDir == "." || currentDir == ".." {
		return api.Error("current_dir is required").ToString()
	}

	if currentDir == "/" {
		currentDir = "" // eliminate double slashes
	}

	dirPath := currentDir + "/" + selectedDirName
	cfmt.Infoln("Deleting directory:", dirPath)
	dirPath = strings.ReplaceAll(dirPath, "//", "/") // remove double slashes
	dirPath = strings.TrimRight(dirPath, "/")        // remove trailing slashes

	if dirPath == "" || dirPath == "/" {
		return api.Error("root directory can not be deleted").ToString()
	}

	cfmt.Infoln("Deleting directory:", dirPath)

	if c.storage == nil {
		return api.Error("Storage not initialized").ToString()
	}

	errDeleted := c.storage.DeleteDirectory(dirPath)

	if errDeleted == nil {
		return api.Success("directory deleted successfully").ToString()
	}

	return api.Error(errDeleted.Error()).ToString()
}

func (c *FileManagerController) fileDeleteAjax(r *http.Request) string {
	selectedFileName := utils.Req(r, "delete_file", "")
	if selectedFileName == "" {
		return api.Error("delete_file is required").ToString()
	}
	currentDir := utils.Req(r, "current_dir", "")
	if currentDir == "" {
		return api.Error("current_dir is required").ToString()
	}

	if currentDir == "/" {
		currentDir = "" // eliminate double slashes
	}

	filePath := currentDir + "/" + selectedFileName

	if c.storage == nil {
		return api.Error("Storage not initialized").ToString()
	}
	errDeleted := c.storage.DeleteFile([]string{filePath})

	if errDeleted == nil {
		return api.Success("file deleted successfully").ToString()
	}

	return api.Error(errDeleted.Error()).ToString()
}

func (c *FileManagerController) fileRenameAjax(r *http.Request) string {
	currentFileName := utils.Req(r, "rename_file", "")
	if currentFileName == "" {
		return api.Error("rename_file is required").ToString()
	}

	newFileName := utils.Req(r, "new_file", "")

	if newFileName == "" {
		return api.Error("new_file is required").ToString()
	}
	currentDir := utils.Req(r, "current_dir", "")

	if currentDir == "" {
		return api.Error("current_dir is required").ToString()
	}

	if currentDir == "/" {
		currentDir = "" // eliminate double slashes
	}

	oldFilePath := currentDir + "/" + currentFileName
	newFilePath := currentDir + "/" + newFileName

	if c.storage == nil {
		return api.Error("Storage not initialized").ToString()
	}

	err := c.storage.Move(oldFilePath, newFilePath)

	if err == nil {
		return api.Success("file renamed successfully").ToString()
	}

	return api.Error(err.Error()).ToString()
}

func (controller *FileManagerController) getMediaManager(r *http.Request) string {
	if controller.storage == nil {
		return api.Error("storage is required").ToString()
	}

	currentDirectory := utils.Req(r, "current_dir", "")
	currentDirectory = strings.Trim(currentDirectory, "/")
	currentDirectory = strings.Trim(currentDirectory, ".")

	parentDirectory := ""
	if currentDirectory != "" {
		parentDirectory = php.Dirname(currentDirectory)
	}

	parentDirectory = strings.Trim(parentDirectory, "/")
	parentDirectory = strings.Trim(parentDirectory, ".")

	if currentDirectory == "" {
		currentDirectory = controller.rootDirPath
	}

	directories, err := controller.storage.Directories(currentDirectory)

	if err != nil {
		return api.Error(err.Error()).ToString()
	}

	files, err := controller.storage.Files(currentDirectory)

	if err != nil {
		return api.Error(err.Error()).ToString()
	}

	directoryList := []FileEntry{}
	for _, dir := range directories {
		size, _ := controller.storage.Size(dir)
		hSize := lo.If(size > 0, controller.HumanFilesize(size)).Else("-")
		modified, _ := controller.storage.LastModified(dir)
		hModified := lo.If(lo.IsEmpty(modified), "-").Else(carbon.CreateFromStdTime(modified).ToDateTimeString())
		directoryList = append(directoryList, FileEntry{
			Path:              dir,
			Name:              php.Basename(dir),
			Size:              size,
			SizeHuman:         hSize,
			LastModified:      modified,
			LastModifiedHuman: hModified,
		})
	}

	fileList := []FileEntry{}
	for _, file := range files {
		size, _ := controller.storage.Size(file)
		hSize := controller.HumanFilesize(size)
		modified, _ := controller.storage.LastModified(file)
		hModified := carbon.CreateFromStdTime(modified).ToDateTimeString()
		url, _ := controller.storage.Url(file)

		fileList = append(fileList, FileEntry{
			Path:              file,
			URL:               url,
			Name:              php.Basename(file),
			Size:              size,
			SizeHuman:         hSize,
			LastModified:      modified,
			LastModifiedHuman: hModified,
		})
	}

	page := controller.uiManager(currentDirectory, parentDirectory, directoryList, fileList)

	if controller.funcLayout != nil {
		style := hb.NewStyleURL(cdn.BootstrapIconsCss_1_10_2()).ToHTML()
		script := hb.NewScriptURL(cdn.Jquery_3_6_4()).ToHTML()
		script += hb.NewScriptURL(cdn.Notify_0_4_2()).ToHTML()
		page = style + script + page
		return controller.funcLayout(page)
	}

	layout := uiLayout("Media Manager", page)
	return layout
}

func (c *FileManagerController) modalFileUpload(currentDirectory string) string {
	url := links.NewAdminLinks().FileManager(map[string]string{})
	return `
<!-- START: Modal Upload File -->
<div class="modal fade" id="ModalUploadFile" role="dialog">
	<div class="modal-dialog" role="document">
		<div class="modal-content">
			<div class="modal-header">
				<h5 class="modal-title" id="myModalLabel">File Upload</h5>
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
			</div>
			<div class="modal-body">
				<form id="FormFileUpload" name="FormFileUpload" method="POST" enctype="multipart/form-data">
				    <input type="hidden" name="action" value="file_upload" />
					<input type="hidden" name="current_dir" value="` + currentDirectory + `" />
					<input type="file" id="file-input" name="upload_file" value="" />
					<input type="hidden" name="_token" value="<?php echo csrf_token(); ?>" />
				</form>
			</div>
			<div class="modal-footer" style="display:block;">
				<button type="button" class="btn btn-secondary float-start" data-bs-dismiss="modal">
					<i class="bi bi-chevron-left"></i>
					Close
				</button>
				<button type="button" class="btn btn-primary float-end" onclick="fileUpload();/*FormFileUpload.submit();*/">
					<i class="bi bi-check"></i>
					Start Upload
				</button>
			</div>
		</div>
	</div>
</div>

<script>
function fileUpload() {
	const file = document.getElementById('file-input').files[0];
	const formData = new FormData();
	formData.append('action', 'file_upload');
	formData.append('current_dir', '` + currentDirectory + `');
	formData.append('upload_file', file);

	try {
		fetch("` + url + `", { method: 'POST', body: formData })
		.then((response) => response.json())
		.then((response) => {
			if (response.status == "success") {
				$.notify(response.message, "success");
			} else {
				$.notify(response.message, "error");	
			}
			$('#ModalUploadFile').modal('hide');
			setTimeout(()=>{
				window.location.href = window.location.href;
			}, 1000)
		})
	} catch (e) {
		$.notify("IO Error", "error");
	}

	// $.post("` + url + `", formData).then((response)=>{
	// 	if (response.status == "success") {
	// 		$.notify(response.message, "success");
	// 	} else {
	// 		$.notify(response.message, "error");	
	// 	}
	// 	$('#ModalUploadFile').modal('hide');
	// 	setTimeout(()=>{
	// 		window.location.href = window.location.href;
	// 	}, 1000)
	// }).fail(()=>{
	// 	$.notify("IO Error", "error");
	// })
}
</script>
<!-- END: Modal Upload File -->
	`
}

func (c *FileManagerController) modalDirectoryCreate(currentDirectory string) string {
	url := links.NewAdminLinks().FileManager(map[string]string{})
	if currentDirectory == "" {
		currentDirectory = "/"
	}
	return `
<div class="modal fade" id="ModalDirectoryCreate" role="dialog">
	<div class="modal-dialog" role="document">
		<div class="modal-content">
			<div class="modal-header">
				<h5 class="modal-title" id="myModalLabel">New Directory</h5>
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
			</div>
			<div class="modal-body">
				<form id="FormDirectoryCreate" name="FormDirectoryCreate"" method="POST">
					<div class="form-group">
						<label>Directory name</label>
						<input type="text" class="form-control" name="create_dir" value="" />
					</div>
					<input type="hidden" name="action" value="` + JSON_ACTION_DIRECTORY_CREATE + `" />
					<input type="hidden" name="current_dir" value="` + currentDirectory + `" />
					<input type="hidden" name="_token" value="<?php echo csrf_token(); ?>" />
				</form>
			</div>
			<div class="modal-footer" style="display:block;">
				<button type="button" class="btn btn-secondary float-start" data-bs-dismiss="modal">
					<i class="bi bi-chevron-left"></i>
					Close
				</button>
				<button type="button" class="btn btn-primary float-end" onclick="directoryCreate();">
					<i class="bi bi-check"></i>
					Create Directory
				</button>
			</div>
		</div>
	</div>
</div>
<script>
	function directoryCreate() {
		$.post("` + url + `", $('#FormDirectoryCreate').serialize()).then((response)=>{
			if (response.status == "success") {
				$.notify(response.message, "success");
			} else {
				$.notify(response.message, "error");	
			}
			$('#ModalDirectoryCreate').modal('hide');
			setTimeout(()=>{
				window.location.href = window.location.href;
			}, 1000)
		}).fail(()=>{
			$.notify("IO Error", "error");
		})
	}
</script>
	`
}

func (c *FileManagerController) modalDirectoryDelete(currentDirectory string) string {
	url := links.NewAdminLinks().FileManager(map[string]string{})
	return `
	<!-- START: Modal Directory Delete -->
	<div class="modal fade" id="ModalDirectoryDelete" role="dialog">
		<div class="modal-dialog" role="document">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="myModalLabel">Confirm Directory Delete</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<p>
						Are you sure you want to delete this folder
						and all the files in it?
					</p>
					<p class="text-danger">
						This operation is final and CANNOT BE undone!
					</p>
					<form id="FormDirectoryDelete" name="FormDirectoryDelete" method="POST">
					    <input type="hidden" name="action" value="` + JSON_ACTION_DIRECTORY_DELETE + `" />
						<input type="hidden" name="current_dir" value="` + currentDirectory + `" />
						<input type="hidden" name="delete_dir" value="" />
						<input type="hidden" name="_token" value="<?php echo csrf_token(); ?>" />
					</form>
				</div>
				<div class="modal-footer" style="display:block;">
					<button type="button" class="btn btn-secondary float-start" data-bs-dismiss="modal">
						<i class="bi bi-chevron-left"></i>
						Close
					</button>
					<button type="button" class="btn btn-danger float-end" onclick="directoryDelete();">
						<i class="bi bi-trash"></i>
						Delete Directory
					</button>
				</div>
			</div>
		</div>
	</div>
	<script>
		function modalDirectoryDeleteShow(directoryName) {
			$('#FormDirectoryDelete input[name="delete_dir"]').val(directoryName);
			const modal = new bootstrap.Modal(document.getElementById('ModalDirectoryDelete'), {})
			modal.show();
		}
		function directoryDelete() {
			$.post("` + url + `", $('#FormDirectoryDelete').serialize()).then((response)=>{
				if (response.status == "success") {
					$.notify(response.message, "success");
				} else {
					$.notify(response.message, "error");	
				}
				$('#FormDirectoryDelete').modal('hide');
				setTimeout(()=>{
					window.location.href = window.location.href;
				}, 1000)
			}).fail(()=>{
				$.notify("IO Error", "error");
			})
		}
	</script>
	<!-- END: Modal Directory Delete -->
	`
}

func (c *FileManagerController) modalFileDelete(currentDirectory string) string {
	url := links.NewAdminLinks().FileManager(map[string]string{})
	return `
	<!-- START: Modal File Delete -->
	<div class="modal fade" id="ModalFileDelete" role="dialog">
		<div class="modal-dialog" role="document">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="myModalLabel">Confirm File Delete</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<p>
						Are you sure you want to delete this file?
					</p>
					<p class="text-danger">
						This operation is final and CANNOT BE undone!
					</p>
					<form id="FormFileDelete" name="FormFileDelete" method="POST">
					    <input type="hidden" name="action" value="` + JSON_ACTION_FILE_DELETE + `" />
						<input type="hidden" name="current_dir" value="` + currentDirectory + `" />
						<input type="hidden" name="delete_file" value="" />
						<input type="hidden" name="_token" value="<?php echo csrf_token(); ?>" />
					</form>
				</div>
				<div class="modal-footer" style="display:block;">
					<button type="button" class="btn btn-secondary float-start" data-bs-dismiss="modal">
						<i class="bi bi-chevron-left"></i>
						Close
					</button>
					<button type="button" class="btn btn-danger float-end" onclick="fileDelete()">
						<i class="bi bi-trash"></i>
						Delete File
					</button>
				</div>
			</div>
		</div>
	</div>
	<script>
		function modalFileDeleteShow(fileName) {
			$('#FormFileDelete input[name="delete_file"]').val(fileName);
			const modal = new bootstrap.Modal(document.getElementById('ModalFileDelete'), {})
			modal.show();
		}
		function fileDelete() {
			$.post("` + url + `", $('#FormFileDelete').serialize()).then((response)=>{
				setTimeout(()=>{
					window.location.href = window.location.href;
				}, 1);
				$('#ModalFileDelete').modal('hide');
				if (response.status == "success") {
					$.notify(response.message, "success");
				} else {
					$.notify(response.message, "error");	
				}
			}).fail(()=>{
				$.notify("IO Error", "error");
			})
		}
	</script>
	<!-- END: Modal File Delete -->
	`
}

func (c *FileManagerController) modalFileRename(currentDirectory string) string {
	url := links.NewAdminLinks().FileManager(map[string]string{})
	return `
<!-- START: Modal File Rename -->
<div class="modal fade" id="ModalFileRename" role="dialog">
	<div class="modal-dialog" role="document">
		<div class="modal-content">
			<div class="modal-header">
				<h5 class="modal-title" id="myModalLabel">File Rename</h5>
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
			</div>
			<div class="modal-body">
				<form id="FormFileRename" name="FormFileRename" action="<?php echo action('Sinevia\Controllers\MediaController@postFileRename'); ?>" method="POST">
					<div class="form-group">
						<label>New Name</label>
						<input name="new_file" value="" class="form-control" />
					</div>
					<input type="hidden" name="action" value="` + JSON_ACTION_FILE_RENAME + `" />
					<input type="hidden" name="current_dir" value="` + currentDirectory + `" />
					<input type="hidden" name="rename_file" value="" />
					<input type="hidden" name="_token" value="<?php echo csrf_token(); ?>" />
				</form>
			</div>
			<div class="modal-footer" style="display:block;">
				<button type="button" class="btn btn-secondary float-start" data-bs-dismiss="modal">
					<i class="bi bi-chevron-left"></i>
					Close
				</button>
				<button type="button" class="btn btn-success float-end" onclick="fileRename()">
					<i class="bi bi-check"></i>
					Rename File
				</button>
			</div>
		</div>
	</div>
</div>
<script>
	function modalFileRenameShow(fileName) {
		$('#FormFileRename input[name="new_file"]').val(fileName);
		$('#FormFileRename input[name="rename_file"]').val(fileName);
		const modal = new bootstrap.Modal(document.getElementById('ModalFileRename'), {})
		modal.show();
	}
	function fileRename() {
		$.post("` + url + `", $('#FormFileRename').serialize()).then((response)=>{
			if (response.status == "success") {
				$.notify(response.message, "success");
			} else {
				$.notify(response.message, "error");
			}
			const modal = new bootstrap.Modal(document.getElementById('ModalFileRename'), {})
			modal.hide();

			setTimeout(()=>{
				//window.location.href = window.location.href;
			}, 1000)
		}).fail(()=>{
			$.notify("IO Error", "error");
		})
	}
</script>
<!-- END: Modal File Rename -->
	`
}

func (c *FileManagerController) modalFileView() string {
	return `
<div class="modal fade" id="ModalFileView" role="dialog">
	<div class="modal-dialog" role="document">
		<div class="modal-content">
			<div class="modal-header">
				<h5 class="modal-title" id="myModalLabel">File View</h5>
				<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
			</div>
			<div class="modal-body" style="text-align:center;">
			    <img id="FileViewImage" src="" class="img-fluid" style="max-height:200px;;" />
			</div>
			<div class="modal-footer" style="display:block;">
				<button type="button" class="btn btn-secondary float-start" data-bs-dismiss="modal">
					<i class="bi bi-chevron-left"></i>
					Close
				</button>
			</div>
		</div>
	</div>
</div>
<script>
	function modalFileViewShow(fileURL) {
		$('#FileViewImage').attr("src", fileURL);
		const modal = new bootstrap.Modal(document.getElementById('ModalFileView'), {})
		modal.show();
	}
</script>
	`
}

func (c *FileManagerController) tableFileList(currentDirectory, parentDirectory string, directoryList, fileList []FileEntry) string {
	table := hb.NewTable().Class("table table-bordered table-striped").Children([]hb.TagInterface{
		hb.NewThead().Children([]hb.TagInterface{
			hb.NewTR().Children([]hb.TagInterface{
				hb.NewTH().Style("width:1px;").Text(""),
				hb.NewTH().Text("Directory/File Name"),
				hb.NewTH().Style("width:100px;").Text("Size"),
				hb.NewTH().Style("width:100px;").Text("Modified"),
				hb.NewTH().Style("width:220px;").Text("Actions"),
			}),
		}),
		hb.NewTbody().
			// Parent DIrectory
			ChildIfF(currentDirectory != "", func() hb.TagInterface {
				parentDirectoryURL := links.NewAdminLinks().FileManager(map[string]string{"current_dir": parentDirectory})

				return hb.NewTR().Children([]hb.TagInterface{
					hb.NewTD().Children([]hb.TagInterface{
						hb.NewI().Class("bi bi-folder").Text(""),
					}),
					hb.NewTD().Children([]hb.TagInterface{
						hb.NewHyperlink().Href(parentDirectoryURL).Children([]hb.TagInterface{
							hb.NewI().Class("bi bi-arrow-90deg-up").Text("").Style("margin-right: 5px;"),
							hb.NewSpan().Text("parent"),
						}),
					}),
					hb.NewTD().Children([]hb.TagInterface{}),
					hb.NewTD().Children([]hb.TagInterface{}),
					hb.NewTD().Children([]hb.TagInterface{}),
				})
			}).
			// Directory List
			ChildIfF(len(directoryList) > 0, func() hb.TagInterface {
				return hb.NewWrap().Children(lo.Map(directoryList, func(dir FileEntry, _ int) hb.TagInterface {
					name := dir.Name
					if dir.Name == "." || dir.Name == ".." {
						return nil
					}
					path := strings.TrimRight(dir.Path, "/")
					pathURL := links.NewAdminLinks().FileManager(map[string]string{"current_dir": path})
					size := dir.SizeHuman

					buttonDelete := hb.NewButton().Class("btn btn-danger btn-sm").OnClick(`modalDirectoryDeleteShow('` + name + `')`).Children([]hb.TagInterface{
						hb.NewI().Class("bi bi-trash").Text("").Style("margin-right: 5px;"),
						hb.NewSpan().Text("Delete"),
					})

					buttonRename := hb.NewButton().Class("btn btn-primary btn-sm").OnClick(`modalFileRenameShow('` + name + `')`).Children([]hb.TagInterface{
						hb.NewI().Class("bi bi-pencil").Text("").Style("margin-right: 5px;"),
						hb.NewSpan().Text("Rename"),
					})

					return hb.NewTR().Children([]hb.TagInterface{
						hb.NewTD().Children([]hb.TagInterface{
							hb.NewI().Class("bi bi-folder").Text(""),
						}),
						hb.NewTD().Children([]hb.TagInterface{
							hb.NewHyperlink().Href(pathURL).Children([]hb.TagInterface{
								hb.NewSpan().Text(name).Style("font-weight: bold;"),
							}),
						}),
						hb.NewTD().Children([]hb.TagInterface{
							hb.NewSpan().Text(size).Style("font-size: 12px;"),
						}),
						hb.NewTD().Children([]hb.TagInterface{
							hb.NewSpan().Text("").Style("font-size: 11px;"),
						}),
						hb.NewTD().Children([]hb.TagInterface{
							buttonRename,
							buttonDelete,
						}),
					})
				}))
			}).
			// File List
			ChildIfF(len(fileList) > 0, func() hb.TagInterface {
				return hb.NewWrap().Children(lo.Map(fileList, func(file FileEntry, _ int) hb.TagInterface {
					buttonDelete := hb.NewButton().Class("btn btn-danger btn-sm").OnClick(`modalFileDeleteShow('` + file.Name + `')`).Children([]hb.TagInterface{
						hb.NewI().Class("bi bi-trash").Text("").Style("margin-right: 5px;"),
						hb.NewSpan().Text("Delete"),
					})

					buttonRename := hb.NewButton().Class("btn btn-primary btn-sm").OnClick(`modalFileRenameShow('` + file.Name + `')`).Children([]hb.TagInterface{
						hb.NewI().Class("bi bi-pencil").Text("").Style("margin-right: 5px;"),
						hb.NewSpan().Text("Rename"),
					})

					buttonView := hb.NewButton().Class("btn btn-info btn-sm").OnClick(`modalFileViewShow('` + file.Name + `')`).Children([]hb.TagInterface{
						hb.NewI().Class("bi bi-eye").Text("").Style("margin-right: 5px;"),
						hb.NewSpan().Text("View"),
					})

					buttonSelect := hb.NewButton().Class("btn btn-success btn-sm .btn-select").OnClick(`fileSelectedUrl('` + file.URL + `')`).Children([]hb.TagInterface{
						hb.NewI().Class("bi bi-chevron-right").Text("").Style("margin-right: 5px;"),
						hb.NewSpan().Text("Select"),
					})

					return hb.NewTR().Children([]hb.TagInterface{
						hb.NewTD().Children([]hb.TagInterface{
							hb.NewI().Class("bi bi-file").Text(""),
						}),
						hb.NewTD().Children([]hb.TagInterface{
							hb.NewSpan().Text(file.Name).Style("font-weight: bold;"),
							hb.NewDiv().
								Children([]hb.TagInterface{
									hb.NewSpan().Text("URL: "),
									hb.NewHyperlink().
										Href(file.URL).
										Target("_blank").
										Children([]hb.TagInterface{
											hb.NewSpan().Text(file.URL),
										}),
								}).
								Style("font-size: 12px;"),
						}),
						hb.NewTD().Children([]hb.TagInterface{
							hb.NewSpan().Text(file.SizeHuman).Style("font-size: 12px;"),
						}),
						hb.NewTD().Children([]hb.TagInterface{
							hb.NewSpan().
								HTML(lo.Substring(file.LastModifiedHuman, 0, 10)).
								Attr("title", file.LastModifiedHuman).
								Style("font-size: 11px;"),
						}),
						hb.NewTD().Children([]hb.TagInterface{
							buttonView,
							buttonRename,
							buttonDelete,
							buttonSelect,
						}),
					})
				}))
			}),
	})
	return table.ToHTML()
}

func uiLayout(title string, content string) string {
	html := `
<!DOCTYPE html>
<html>
    <head>
        <title>` + title + ` - Media Manager</title>
        <link href="` + cdn.BootstrapIconsCss_1_10_2() + `" rel="stylesheet" type="text/css" />
		<link href="` + cdn.BootstrapCss_5_2_3() + `" rel="stylesheet" type="text/css" />
        <script src="` + cdn.Jquery_3_6_4() + `"></script>
        <script src="` + cdn.BootstrapJs_5_2_3() + `"></script>
		<script src="` + cdn.Notify_0_4_2() + `"></script>
        <style>
            html,body{
                padding-top: 40px;
            }
        </style>
    </head>
    <body>
        <!-- Fixed navbar -->
        <nav class="navbar navbar-expand-lg bg-light fixed-top"  data-bs-theme="dark">
            <div class="container">
				<a class="navbar-brand" href="#">
					Media Manager
				</a>
				<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarTogglerDemo01" aria-controls="navbarTogglerDemo01" aria-expanded="false" aria-label="Toggle navigation">
					<span class="navbar-toggler-icon"></span>
				</button>
            </div>
        </nav>
        <div class="container">` + content + ` </div>
    </body>
</html>
	`

	return html
}

func (c *FileManagerController) uiManager(currentDirectory, parentDirectory string, directoryList, fileList []FileEntry) string {
	buttonUpload := hb.NewButton().
		Class("btn btn-secondary float-end").
		Data("bs-toggle", "modal").
		Data("bs-target", "#ModalUploadFile").
		// OnClick(`modalFileUploadShow()`).
		Child(hb.NewI().Class("bi bi-upload").Style("margin-right: 5px;")).
		HTML("Upload file")

	buttonDirectoryCreate := hb.NewButton().
		Class("btn btn-info float-end me-2").
		Data("bs-toggle", "modal").
		Data("bs-target", "#ModalDirectoryCreate").
		// OnClick(`modalDirectoryCreateShow()`).
		Child(hb.NewI().Class("bi bi-plus-circle").Style("margin-right: 5px;")).
		HTML("New directory")

	title := hb.NewHeading3().
		HTML("Media Manager").
		Child(buttonUpload).
		Child(buttonDirectoryCreate)

	script := hb.NewScript(`
$('.btn-select').hide();
	
var openerArgs = {};

function fileSelectedUrl(selectedFileUrl) {
	if (window.opener === null) {
		return true;
	}
	window.opener.postMessage({msg: 'media-manager-file-selected', url: selectedFileUrl, args: openerArgs}, '*');
	window.close();
}

function messageReceived(event) {
	var data = event.data;
	openerArgs = data;
	console.log(data);
	$('.btn-select').show();
}

window.addEventListener("message", messageReceived, false);

if (window.opener !== null) {
	window.opener.postMessage({msg: 'media-manager-loaded'}, '*');
}
	`)

	html := title.ToHTML() + `
	` + c.tableFileList(currentDirectory, parentDirectory, directoryList, fileList) + `
	` + c.modalDirectoryCreate(currentDirectory) + `
	` + c.modalDirectoryDelete(currentDirectory) + `
	` + c.modalFileDelete(currentDirectory) + `
	` + c.modalFileRename(currentDirectory) + `
	` + c.modalFileUpload(currentDirectory) + `
	` + c.modalFileView() +
		script.ToHTML()

	return html
}

func (c *FileManagerController) HumanFilesize(size int64) string {
	const unit = 1000
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(size)/float64(div), "kMGTPE"[exp])
}
