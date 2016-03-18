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

// Default values
const (
	port      = 8080      // default port number
	uploadDir = "uploads" // file upload directory
)

// Request parameters
const (
	paramUuid = "qquuid" // uuid
	paramFile = "qqfile" // file name
)

// Chunked request parameters
const (
	paramPartIndex       = "qqpartindex"      // part index
	paramPartBytesOffset = "qqpartbyteoffset" // part byte offset
	paramTotalFileSize   = "qqtotalfilesize"  // total file size
	paramTotalParts      = "qqtotalparts"     // total parts
	paramFileName        = "qqfilename"       // file name for chunked requests
	paramChunkSize       = "qqchunksize"      // size of the chunks
)

type UploadResponse struct {
	Success      bool   `json:"success"`
	Error        string `json:"error,omitempty"`
	PreventRetry bool   `json:"preventRetry"`
}

func main() {
	log.Printf("Initiating server on port [%d]\n", port)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/upload", UploadHandler)
	http.Handle("/upload/", http.StripPrefix("/upload/", http.HandlerFunc(UploadHandler)))
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

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		upload(w, req)
		return
	case http.MethodDelete:
		delete(w, req)
		return
	}
	errorMsg := fmt.Sprintf("Method [%s] is not supported:", req.Method)
	http.Error(w, errorMsg, http.StatusMethodNotAllowed)
}

func upload(w http.ResponseWriter, req *http.Request) {
	uuid := req.FormValue(paramUuid)
	log.Printf("Starting upload handling of request with uuid of [%s]\n", uuid)
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

	var filename string
	partIndex := req.FormValue(paramPartIndex)
	if len(partIndex) == 0 {
		filename = fmt.Sprintf("%s/%s", fileDir, headers.Filename)

	} else {
		filename = fmt.Sprintf("%s/%s_%05s", fileDir, uuid, partIndex)
	}
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

func delete(w http.ResponseWriter, req *http.Request) {
	log.Printf("Delete request received for uuid [%s]", req.URL.Path)
	err := os.RemoveAll(uploadDir + "/" + req.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}
