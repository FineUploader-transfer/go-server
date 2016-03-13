// An upload server for fineuploader.com javascript upload library
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	// default port number
	port = 8080
	// file upload directory
	uploadDir = "uploads"
	// fineuploader uuid param name
	paramUuid = "qquuid"
	// fineuploader file param name
	paramFile = "qqfile"
)

type UploadResponse struct {
	Success      bool   `json:"success"`
	Error        string `json:"error,omitempty"`
	PreventRetry bool   `json:"preventRetry"`
}

func main() {
	log.Printf("Initiating server on port [%d]\n", port)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/upload", Upload)
	http.Handle("/delete/", http.StripPrefix("/delete/", http.HandlerFunc(Delete)))
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil))
}

func writeUploadResponse(w http.ResponseWriter, err error) {
	uploadResponse := new(UploadResponse)
	if err != nil {
		uploadResponse.Error = err.Error()
	} else {
		uploadResponse.Success = true
	}
	w.Header().Set("Content-Type", "text/plain")
	json.NewEncoder(w).Encode(uploadResponse)
}

func Upload(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Not supported http method", http.StatusMethodNotAllowed)
		return
	}
	uuid := req.FormValue(paramUuid)
	log.Printf("Trying to save file with uuid of [%s]\n", uuid)
	file, headers, err := req.FormFile(paramFile)
	if err != nil {
		writeUploadResponse(w, err)
		return
	}

	fileDir := fmt.Sprintf("%s/%s", uploadDir, uuid)
	if err := os.MkdirAll(fileDir, 0777); err != nil {
		writeUploadResponse(w, err)
		return
	}

	filename := fmt.Sprintf("%s/%s", fileDir, headers.Filename)
	outfile, err := os.Create(filename)
	if err != nil {
		writeUploadResponse(w, err)
		return
	}
	defer outfile.Close()

	_, err = io.Copy(outfile, file)
	if err != nil {
		writeUploadResponse(w, err)
		return
	}

	writeUploadResponse(w, nil)
}

func Delete(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Not supported http method", http.StatusMethodNotAllowed)
		return
	}
	log.Printf("Delete request received for uuid [%s]", req.URL.Path)
	err := os.RemoveAll(uploadDir + "/" + req.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}
