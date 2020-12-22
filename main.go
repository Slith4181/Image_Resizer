package main

import (
	"fmt"
	"net/http"
	"time"
)

func index(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "resources/html/index.html")
}

func uploadImage(wr http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1024 << 15)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("\nFile Name: %+v", handler.Filename)
	fmt.Printf("\nFile Name: %+v", handler.Size)
	fmt.Printf("\nMIME Name: %+v", handler.Header)

	name := "upload-*.png"

	/* tempFile, err := ioutil.TempFile("resources/images", name)
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileByte) */

	http.ServeContent(wr, r, name, time.Now(), file)

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadImage)
	http.ListenAndServe(":8080", nil)
}
