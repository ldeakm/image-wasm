package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"io"
	"strings"
	"syscall/js"
	"time"

	"github.com/nfnt/resize"
)

var done chan bool
var imgEle js.Value
var start time.Time

const (
	FRAMEDELAY int = 1000
)

func main() {
	// Something to block to keep the wasm from existing
	done = make(chan bool)

	// Create a call back for the camera
	camCallback := js.NewCallback(baseToImage)
	defer camCallback.Release()

	// Call the function in the js file to start
	// the camera capture
	cap := js.Global().Get("startCapture")
	cap.Invoke(camCallback, 5)

	// Create the elements needed for the modified image
	imgEle = js.Global().Get("document").Call("createElement", "img")
	js.Global().Get("document").Get("body").Call("appendChild", imgEle)

	// Block exit
	<-done
}

func outPage(img2 image.Image) {

	// Pulled in external package for a wuick resize
	img := resize.Resize(160, 0, img2, resize.Lanczos3)

	// prep write out
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Println(err.Error())
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	// write out the time it took to process the image
	fmt.Println(time.Since(start))

	// set the image tag to the precessed image
	imgEle.Set("src", "data:image/png;base64,"+imgBase64Str)

}

func baseToImage(args []js.Value) {
	// timer to get a feel for timing
	start = time.Now()
	p := strings.Split(args[0].String(), ",")
	img, err := png.Decode(gpng(p[1]))
	if err != nil {
		fmt.Println(err.Error())
	}
	outPage(img)
}

func gpng(gopher string) io.Reader {
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(gopher))
}
