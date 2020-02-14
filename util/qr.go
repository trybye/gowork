package util


import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func Qrcode(content string) {
	// Create the barcode
	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 100, 100)

	// create the output file
	file, _ := os.Create("/data/tools/app/pay/img/"+content+".png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
