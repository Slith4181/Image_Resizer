package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nfnt/resize"
)

func index(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "resources/html/index.html")
}

func uploadAnPicture(wr http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(8 << 20)

	picture, handler, err := r.FormFile("pic")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer picture.Close()

	fmt.Printf("\nFile Name: %+v", handler.Filename)
	fmt.Printf("\nFile Name: %+v", handler.Size)
	fmt.Printf("\nMIME Name: %+v", handler.Header)

	buff := make([]byte, 512)

	if _, err = picture.Read(buff); err != nil {
		fmt.Println(err)
		return
	}
	//working code

	buffer := new(bytes.Buffer)

	Decode_Image, _, err := image.Decode(bytes.NewReader(buff))
	if err != nil {
		log.Fatal(err)
	}
	Resized_Image := resize.Resize(500, 0, Decode_Image, resize.Lanczos3) //Ресайз декодированного изображения

	files, err := os.Create(handler.Filename)
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(buffer, Resized_Image)

	tempFile, err := ioutil.TempFile("File_storage", handler.Filename)
	defer tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(tempFile, files); err != nil {
		log.Fatal(err)
		return
	}
	//working code
	if http.DetectContentType(buff) == "image/png" || http.DetectContentType(buff) == "image/img" || http.DetectContentType(buff) == "image/jpeg" {
		http.ServeContent(wr, r, handler.Filename, time.Now(), picture)
	} else {
		http.Error(wr, "Invalid file format", http.StatusBadRequest)
		return
	} // Handle error
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadAnPicture)
	http.ListenAndServe(":8080", nil)

}
