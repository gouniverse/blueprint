package shared

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"project/config"
	"project/internal/helpers"
	"project/internal/links"
	"project/internal/resources"
	"time"

	"strings"

	"github.com/disintegration/imaging"
	"github.com/go-chi/chi/v5"
	"github.com/gouniverse/router"
	"github.com/gouniverse/utils"
	"github.com/mingrammer/cfmt"
	"github.com/samber/lo"
)

// == CONTROLLER ==============================================================

type thumbnailController struct{}

// == CONSTRUCTOR =============================================================

func NewThumbController() *thumbnailController {
	return &thumbnailController{}
}

var _ router.HTMLControllerInterface = (*thumbnailController)(nil)

// ThumbnailHandler
// ================================================================
// Resizes local images to the specified width and height
// ================================================================
// Path
// /th/EXT/WIDTHxHEIGHT/QUALITY/path
// Example:
// /th/jpg/1920x0/70/images/backgrounds/pexels-pixabay-531756.jpg
// ================================================================
func (controller *thumbnailController) Handler(w http.ResponseWriter, r *http.Request) string {
	data, errorMessage := controller.prepareData(r)

	if errorMessage != "" {
		return errorMessage
	}

	cacheKey := utils.StrToMD5Hash(fmt.Sprint(data.path, data.extension, data.width, "x", data.height, data.quality))

	if config.CacheFile.Contains(cacheKey) {
		thumb, err := config.CacheFile.Fetch(cacheKey)

		if err == nil {
			controller.setHeaders(w, data.extension)
			return thumb
		}
	}

	thumb, errorMessage := controller.generateThumb(data)

	if errorMessage != "" {
		return errorMessage
	}

	err := config.CacheFile.Save(cacheKey, thumb, 5*time.Minute) // cache for 5 minutes

	if err != nil {
		cfmt.Errorln("Error at thumbnailController > CacheFile.Save", "error", err.Error())
	}

	controller.setHeaders(w, data.extension)
	return thumb
}

func (controller *thumbnailController) setHeaders(w http.ResponseWriter, fileExtension string) {
	w.Header().Set("Content-Type", lo.
		If(fileExtension == "jpg", "image/jpeg").
		ElseIf(fileExtension == "jpeg", "image/jpeg").
		ElseIf(fileExtension == "png", "image/png").
		ElseIf(fileExtension == "gif", "image/gif").
		Else(""))

	w.Header().Set("Cache-Control", "max-age=604800") // cache for SEO
}

func (controller *thumbnailController) prepareData(r *http.Request) (data thumbnailControllerData, errorMessage string) {
	data.extension = chi.URLParam(r, "extension")
	size := chi.URLParam(r, "size")
	quality := chi.URLParam(r, "quality")
	data.path = chi.URLParam(r, "*")
	data.isURL = false

	///cfmt.Infoln("====================================")
	//cfmt.Infoln("EXTENSION: ", extension)
	//cfmt.Infoln("SIZE: ", size)
	//cfmt.Infoln("QUALITY: ", quality)
	//cfmt.Infoln("PATH: ", path)
	//cfmt.Infoln("====================================")

	if data.extension == "" {
		return data, "image extension is missing"
	}

	if size == "" {
		return data, "size is missing"
	}

	if quality == "" {
		return data, "quality is missing"
	}

	if data.path == "" {
		return data, "path is missing"
	}

	if strings.HasPrefix(data.path, "http/") || strings.HasPrefix(data.path, "https/") {
		data.isURL = true
		data.path = strings.ReplaceAll(data.path, "https/", "https://")
		data.path = strings.ReplaceAll(data.path, "http/", "http://")
	}

	if strings.HasPrefix(data.path, "files/") {
		data.path = links.URL(data.path, nil)
		data.isURL = true
	}

	widthStr := ""
	heightStr := ""
	if strings.Contains(size, "x") {
		splits := strings.Split(size, "x")
		widthStr = splits[0]
		heightStr = splits[1]
	} else {
		widthStr = size
	}

	widthInt, _ := utils.StrToInt64(widthStr)
	heightInt, _ := utils.StrToInt64(heightStr)
	qualityInt, _ := utils.StrToInt64(quality)

	data.width = widthInt
	data.height = heightInt
	data.quality = qualityInt

	return data, errorMessage
}

func (controller *thumbnailController) generateThumb(data thumbnailControllerData) (content string, errorMessage string) {
	ext := imaging.JPEG

	if data.extension == "gif" {
		ext = imaging.GIF
	}

	if data.extension == "png" {
		ext = imaging.PNG
	}

	// cfmt.Infoln("EXTENSION: ", ext)
	// cfmt.Infoln("WIDTH: ", data.width)
	// cfmt.Infoln("HEIGHT: ", data.height)
	// cfmt.Infoln("QUALITY: ", data.quality)
	// cfmt.Infoln("PATH: ", data.path)

	var err error
	var imgBytes []byte

	if data.isURL {
		//imgBytes = controller.toBytes(data.path)
		imgBytes, err = controller.urlToBytes(data.path)

		if err != nil {
			config.Logger.Error("Error at thumbnailController > generateThumb > from URL", "error", err.Error())
			return "", err.Error()
		}
	} else {
		var err error
		imgBytes, err = resources.ToBytes(data.path)

		if err != nil {
			config.Logger.Error("Error at thumbnailController > generateThumb > from RESOURCE", "error", err.Error())
			return "", err.Error()
		}
	}

	imgBytesResized, err := helpers.ImageResize(imgBytes, int(data.width), int(data.height), ext)

	if err != nil {
		config.Logger.Error("Error at thumbnailController > generateThumb", "error", err.Error())
		return "", err.Error()
	}

	return string(imgBytesResized), ""
}

// func (controller *thumbnailController) thumbnailProcess(w http.ResponseWriter, r *http.Request) (content string, errorMessage string) {
// 	thumbnail, errorMessage := controller.thumbnailValidate(r)

// 	if errorMessage != "" {
// 		config.LogStore.ErrorWithContext("Error at thumbnailController", errorMessage)
// 		return "", errorMessage
// 	}

// 	ext := imaging.JPEG

// 	if thumbnail.Extension == "gif" {
// 		ext = imaging.GIF
// 	}
// 	if thumbnail.Extension == "png" {
// 		ext = imaging.PNG
// 	}

// 	var imgBytes []byte

// 	if thumbnail.IsURL {
// 		imgBytes = controller.toBytes(thumbnail.Path)
// 	} else {
// 		var err error
// 		imgBytes, err = media.ToBytes(thumbnail.Path)

// 		if err != nil {
// 			config.LogStore.ErrorWithContext("Error at thumbnailController", err.Error())
// 			return "", err.Error()
// 		}
// 	}

// 	imgBytesResized, err := helpers.ImageResize(imgBytes, int(thumbnail.Width), int(thumbnail.Height), ext)

// 	if err != nil {
// 		config.LogStore.ErrorWithContext("Error at thumbnailController", err.Error())
// 		return "", err.Error()
// 	}

// 	return string(imgBytesResized), ""
// }

// funnelValidate validates that the funnel exists,
// and that the user has access to it
// func (controller *thumbnailController) thumbnailValidate(r *http.Request) (th *Thumbnail, errorMessage string) {
// 	extension := chi.URLParam(r, "extension")
// 	size := chi.URLParam(r, "size")
// 	quality := chi.URLParam(r, "quality")
// 	path := chi.URLParam(r, "*")
// 	isURL := false

// 	///cfmt.Infoln("====================================")
// 	//cfmt.Infoln("EXTENSION: ", extension)
// 	//cfmt.Infoln("SIZE: ", size)
// 	//cfmt.Infoln("QUALITY: ", quality)
// 	//cfmt.Infoln("PATH: ", path)
// 	//cfmt.Infoln("====================================")

// 	if extension == "" {
// 		return nil, "image extension is missing"
// 	}

// 	if size == "" {
// 		return nil, "size is missing"
// 	}

// 	if quality == "" {
// 		return nil, "quantity is missing"
// 	}

// 	if path == "" {
// 		return nil, "path is missing"
// 	}

// 	if strings.HasPrefix(path, "http/") || strings.HasPrefix(path, "https/") {
// 		isURL = true
// 		url := strings.ReplaceAll(path, "https/", "https://")
// 		url = strings.ReplaceAll(url, "http/", "http://")

// 		cfmt.Infoln("URL: ", url)

// 		localFilePath := os.TempDir() + "/" + uid.HumanUid() + "." + extension

// 		cfmt.Infoln("LOCAL PATH: ", localFilePath)

// 		errDownload := helpers.DownloadFile(localFilePath, url)
// 		if errDownload != nil {
// 			return nil, "download failed"
// 		}

// 		path = localFilePath
// 	}

// 	widthStr := ""
// 	heightStr := ""
// 	if strings.Contains(size, "x") {
// 		splits := strings.Split(size, "x")
// 		widthStr = splits[0]
// 		heightStr = splits[1]
// 	} else {
// 		widthStr = size
// 	}

// 	widthInt, _ := utils.StrToInt64(widthStr)
// 	heightInt, _ := utils.StrToInt64(heightStr)
// 	qualityInt, _ := utils.StrToInt64(quality)

// 	th = &Thumbnail{
// 		IsURL:     isURL,
// 		Extension: extension,
// 		Path:      path,
// 		Width:     widthInt,
// 		Height:    heightInt,
// 		Quality:   qualityInt,
// 	}

// 	return th, ""
// }

func (controller *thumbnailController) urlToBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Url: " + url + " NOT FOUND")
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Url: " + url + " NOT FOUND")
		return nil, err
	}
	return body, nil
}

func (controller *thumbnailController) toBytes(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Path: " + path + " NOT FOUND")
		return nil, err
	}
	return bytes, nil
}

type thumbnailControllerData struct {
	extension string
	width     int64
	height    int64
	quality   int64
	path      string
	isURL     bool
}
