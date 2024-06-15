package shared

import (
	"net/http"
	"project/config"
	"strings"

	"github.com/gouniverse/utils"
	"github.com/samber/lo"
)

type mediaController struct {
}

func NewMediaController() *mediaController {
	return &mediaController{}
}

func (c *mediaController) Handler(w http.ResponseWriter, r *http.Request) string {

	filePath := lo.IfF(strings.HasPrefix(r.URL.Path, "/files"), func() string { return utils.StrRightFrom(r.URL.Path, "/files") }).
		ElseIfF(strings.HasPrefix(r.URL.Path, "/file"), func() string { return utils.StrRightFrom(r.URL.Path, "/file") }).
		ElseIfF(strings.HasPrefix(r.URL.Path, "/media"), func() string { return utils.StrRightFrom(r.URL.Path, "/media") }).
		Else(r.URL.Path)

	exists, err := config.SqlFileStorage.Exists(filePath)

	if err != nil {
		return err.Error()
	}

	if !exists {
		return "File not found"
	}

	content, err := config.SqlFileStorage.ReadFile(filePath)

	if err != nil {
		return err.Error()
	}

	extension := c.findExtension(filePath)

	if extension == "" {
		return "File not found"
	}

	if extension == "html" {
		w.Header().Set("Content-Type", "text/html")
	} else if extension == "css" {
		w.Header().Set("Content-Type", "text/css")
	} else if extension == "js" {
		w.Header().Set("Content-Type", "application/javascript")
	} else if extension == "json" {
		w.Header().Set("Content-Type", "application/json")
	} else if extension == "png" {
		w.Header().Set("Content-Type", "image/png")
	} else if extension == "jpg" || extension == "jpeg" {
		w.Header().Set("Content-Type", "image/jpeg")
	} else if extension == "gif" {
		w.Header().Set("Content-Type", "image/gif")
	} else if extension == "svg" {
		w.Header().Set("Content-Type", "image/svg+xml")
	} else if extension == "ico" {
		w.Header().Set("Content-Type", "image/x-icon")
	} else if extension == "pdf" {
		w.Header().Set("Content-Type", "application/pdf")
	} else if extension == "zip" {
		w.Header().Set("Content-Type", "application/zip")
	} else if extension == "mp3" {
		w.Header().Set("Content-Type", "audio/mpeg")
	} else {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename="+r.URL.Path)
		w.Header().Set("Content-Length", string(len(content)))
	}

	w.Write(content)

	return ""
}

func (controller mediaController) findFileName(path string) string {
	uriParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(uriParts) < 1 {
		return ""
	}

	return uriParts[len(uriParts)-1]
}

// findExtension finds the file extension from a path.
//
// Parameter(s):
//   - path string - the path
//
// Return type(s):
//   - string - the file extension
func (controller mediaController) findExtension(path string) string {
	fileName := controller.findFileName(path)

	if fileName == "" {
		return ""
	}

	nameParts := strings.Split(fileName, ".")

	if len(nameParts) < 2 {
		return ""
	}

	return nameParts[1]
}
