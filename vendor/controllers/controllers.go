package controllers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

// Index Index Page
func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

// Upload This controller works for each file in the upload list.
func Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("files[]")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", fileHeader.Header)
	f, err := os.OpenFile("./uploads/"+fileHeader.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
