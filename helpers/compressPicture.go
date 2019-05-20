package helpers

import (
	"bytes"
	"github.com/astaxie/beego/context"
	"golang_api/src/github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
)

func CompressPicture(ctx *context.Context) {
	var buf bytes.Buffer
	file, err := os.Open("resource/images/6Co6W3DogOqW6TxwDp8Vb.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	m := resize.Resize(50, 50, img, resize.Lanczos3)

	// write new image to file
	jpeg.Encode(&buf, m, nil)
	file.Close()
	data := buf.Bytes()
	ctx.Output.Body(data)
}