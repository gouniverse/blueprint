package helpers

import (
	"bytes"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gouniverse/utils"
)

func UrlToQrBase64String(url string, width int, height int) string {
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, width, height)

	var buffer bytes.Buffer
	png.Encode(&buffer, qrCode)

	qrCodeImageString := utils.BytesToBase64Url(buffer.Bytes())
	return qrCodeImageString
}
