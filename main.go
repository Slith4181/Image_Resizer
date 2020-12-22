package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func index(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "resources/html/index.html")
}

func displayAnImage(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "resources/images/cat.jpg")
}

func uploadImage(w http.ResponseWriter, r *http.Request) {

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

	tempFile, err := ioutil.TempFile("resources/images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileByte)

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadImage)
	http.ListenAndServe(":8080", nil)
}
