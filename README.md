# Go (a.k.a. Golang) Server-Side Example for the Widen Fine Uploader Javascript Library #

This repository contains a [**Golang**](https://golang.org/) server-side example for users of [**Widen's Fine Uploader javascript library**](http://fineuploader.com/).  

### This server supports

* [File chunking](http://docs.fineuploader.com/branch/master/features/chunking.html)
* [Concurrent uploading](http://docs.fineuploader.com/branch/master/features/concurrent-chunking.html)
* [Delete uploaded files](http://docs.fineuploader.com/branch/master/features/delete.html)
* [Resume uploads](http://docs.fineuploader.com/branch/master/features/resume.html)

### Requirements

Go 1.6.1 (should work with previous versions with minor tweaks)
No additional dependencies

## Getting Started

### Server installation

Run `go get` pointing to the repository,
```bash
$ go get github.com/FineUploader/fineuploader-go-server
```

### Compile and install

```bash
$ cd $GOPATH/src/github.com/FineUploader/fineuploader-go-server
$ go install
```

### Run the server

```bash
$ $GOPATH/bin/fineuploader-go-server
```

#### Server start up flags

You can configure the server on start up with the below flags,

Flag | Default value | Description
-----| ------------- | ------------
p | 8080 | Listening port
d | uploads | Root upload directory

Example:
```bash
$ $GOPATH/bin/fineuploader-go-server -p 9000 -d myuploaddir
```

### Server endpoints
Method | Endpoint | Usage
-------|----------|-------
POST|/upload|Upload file end point. Will create `<root_upload_directory>/qquuid` directory and store the received file inside. Refer to [request.endpoint](http://docs.fineuploader.com/branch/master/api/options.html#request.endpoint)
DELETE|/upload|Deletes the uploaded file based on the `qquuid` parameter. Refer to [deleteFile.endpoint](http://docs.fineuploader.com/branch/master/api/options.html#deleteFile.endpoint)
POST|/chunksdone|Builds original file based on received chunks for the received `qquuid` parameter. Refer to [chunking.success.endpoint](http://docs.fineuploader.com/branch/master/api/options.html#chunking.success.endpoint)


### Fine Uploader library

Go to the [Download](http://fineuploader.com/customize.html) section of the Fine Uploader website for instructions on how to get the library.

Go to the [Demos](http://fineuploader.com/demos.html) section for Fine Uploader configuration examples.

### License ###
This project is licensed under the terms of the MIT license.
