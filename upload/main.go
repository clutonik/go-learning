package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Code to demonstrate file upload.")
	http.HandleFunc("/", fileUpload)
	// Start listner using defaultMuxer
	http.ListenAndServe(":8080", nil)
}

// fileUpload demonstrates how to upload a file using HTML form
func fileUpload(w http.ResponseWriter, r *http.Request) {
	// Check if HTTP Request Method is POST (uses const. from http package)
	if r.Method == http.MethodPost {
		// Get file uploaded through HTML Form using http.Request.FormFile method
		f, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		defer f.Close()

		// Read contents of file using ioutil.ReadAll which returns slice of bytes and error
		fileContents, err := ioutil.ReadAll(f)
		if err != nil {
			// Throw internal server error if there was an error while reading file contents
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "File contents: %v", string(fileContents))
	} else {
		// If HTTP Request Method is not POST
		w.Header().Set("Content-Type", "text/html, charset=utf-8")
		io.WriteString(w, `
		<form method="POST" enctype="multipart/form-data">
		<input type="file" name="file">
		<input type="submit">
		</form>
		`)
	}
}
