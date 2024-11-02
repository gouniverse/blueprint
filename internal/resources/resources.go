package resources

import (
	"bytes"
	"embed"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
)

//go:embed *
var files embed.FS

func ToBytes(path string) ([]byte, error) {
	return files.ReadFile(path)
}

func ToString(path string) (string, error) {
	bytes, err := ToBytes(path)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func ResourceExists(path string) bool {
	_, err := files.ReadFile(path)
	return err == nil
}

func Resource(path string) string {
	str, err := files.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(str)
}

func ResourceWithParams(path string, params map[string]string) string {
	parsed := template.Must(template.ParseFS(files, path))
	var tpl bytes.Buffer
	if err := parsed.Execute(&tpl, params); err != nil {
		log.Println(err)
		return ""
	}

	return tpl.String()
}

func ImageToBase64String(path string) string {
	data, _ := files.ReadFile(path)

	mimeType := http.DetectContentType(data)

	base64Encoding := ""

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/bmp":
		base64Encoding += "data:image/bmp;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/webp":
		base64Encoding += "data:image/webp;base64,"
	}

	base64Encoding += base64.URLEncoding.EncodeToString([]byte(data))
	return base64Encoding
}
