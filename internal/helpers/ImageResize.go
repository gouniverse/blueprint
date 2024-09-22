package helpers

import (
	"bytes"
	"encoding/base64"
	"image"
	"net/http"

	"github.com/disintegration/imaging"
)

func ImageBlur(content []byte, blur float64, format imaging.Format) ([]byte, error) {
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	// Create a blurred version of the image.
	dstImage := imaging.Blur(srcImage, blur)

	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	return buffer.Bytes(), errImageEncode
}

func ImageGrayscale(content []byte, format imaging.Format) ([]byte, error) {
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	dstImage := imaging.Grayscale(srcImage)

	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	return buffer.Bytes(), errImageEncode
}

func ImageResize(content []byte, width, height int, format imaging.Format) ([]byte, error) {
	srcImage, _, errImageDecode := image.Decode(bytes.NewReader(content))

	if errImageDecode != nil {
		return nil, errImageDecode
	}

	dstImage := imaging.Resize(srcImage, width, height, imaging.Lanczos)

	var buffer bytes.Buffer
	errImageEncode := imaging.Encode(&buffer, dstImage, format)

	if errImageEncode != nil {
		return nil, errImageEncode
	}

	return buffer.Bytes(), errImageEncode
}

func ImageToBase64String(imgBytes []byte) string {
	mimeType := http.DetectContentType(imgBytes)

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

	// result := tpl.String()
	base64Encoding += base64Encode(imgBytes)
	return base64Encoding
}

func base64Encode(src []byte) string {
	return base64.RawStdEncoding.EncodeToString(src)
	// return base64.URLEncoding.EncodeToString(src)
}
