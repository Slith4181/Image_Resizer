package main

import (
	"fmt"
	"image/png"
	"net/http"
	"time"

	"github.com/nfnt/resize"
)

func index(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "resources/html/index.html")
}

func uploadAnImage(wr http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(8 << 20)

	image, handler, err := r.FormFile("pic")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer image.Close()

	fmt.Printf("\nFile Name: %+v", handler.Filename)
	fmt.Printf("\nFile Name: %+v", handler.Size)
	fmt.Printf("\nMIME Name: %+v", handler.Header)

	buff := make([]byte, 512)

	if _, err = image.Read(buff); err != nil {
		fmt.Println(err)
		return
	}

	inputImage, err := png.Decode(image)
	imgResized := resize.Resize(1366, 768, inputImage, resize.MitchellNetravali)

	fmt.Println(http.DetectContentType(buff))

	if http.DetectContentType(buff) == "image/png" || http.DetectContentType(buff) == "image/img" || http.DetectContentType(buff) == "image/jpeg" {
		http.ServeContent(wr, r, handler.Filename, time.Now(), image)
	} else {
		http.Error(wr, "Invalid file format", http.StatusBadRequest)
		return
	} // Handle error
	png.Encode(wr, imgResized)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadAnImage)
	http.ListenAndServe(":8080", nil)

}
