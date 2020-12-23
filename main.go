package main

import (
	"fmt"
	"net/http"
	"time"
)

func index(wr http.ResponseWriter, r *http.Request) {
	http.ServeFile(wr, r, "resources/html/index.html")
}

func uploadAnImage(wr http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(8 << 20)

	file, handler, err := r.FormFile("pic")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("\nFile Name: %+v", handler.Filename)
	fmt.Printf("\nFile Name: %+v", handler.Size)
	fmt.Printf("\nMIME Name: %+v", handler.Header)

	buff := make([]byte, 512)

	if _, err = file.Read(buff); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(http.DetectContentType(buff))
	if http.DetectContentType(buff) == "image/png" || http.DetectContentType(buff) == "image/img" || http.DetectContentType(buff) == "image/jpeg" {
		http.ServeContent(wr, r, handler.Filename, time.Now(), file)
	} else {
		http.Error(wr, "Invalid file format", http.StatusBadRequest)
		return
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadAnImage)
	http.ListenAndServe(":8080", nil)

}
