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

func Upload(res http.ResponseWriter, req *http.Request) {
	uplRes := new(UploadResponse)

	uuid := req.FormValue(paramUuid)
	log.Printf("Trying to save file with uuid of [%s]\n", uuid)
	file, headers, err := req.FormFile(paramFile)
	if err != nil {
		uplRes.Error = err.Error()
		json.NewEncoder(res).Encode(uplRes)
		return
	}

	fileDir := fmt.Sprintf("%s/%s", uploadDir, uuid)
	os.MkdirAll(fileDir, 0777)
	filename := fmt.Sprintf("%s/%s", fileDir, headers.Filename)

	outfile, err := os.Create(filename)
	defer outfile.Close()

	_, err = io.Copy(outfile, file)
	if err != nil {
		uplRes.Error = err.Error()
		json.NewEncoder(res).Encode(uplRes)
		return
	}

	uplRes.Success = true
	json.NewEncoder(res).Encode(uplRes)
}

func Delete(res http.ResponseWriter, req *http.Request) {
	log.Printf("Delete request received for uuid [%s]", req.URL.Path)
	err := os.RemoveAll(uploadDir + "/" + req.URL.Path)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusOK)

}
