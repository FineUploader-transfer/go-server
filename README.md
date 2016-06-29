# Golang Server-Side Example for the Widen Fine Uploader Javascript Library #

This repository contains a [**Golang**](https://golang.org/) server-side example for users of [**Widen's Fine Uploader javascript library**](http://fineuploader.com/).  

##### This server supports

* File chunking
* Concurrent uploading
* Delete uploaded files
* Pause / Resume uploads

##### Requirements

This server example only uses the Golang standard library so it has no dependencies, you only need to have Go installed in your machine. I developed this using Go 1.6 but it should run on previous versions.

##### Get the code

###### Setup your project structure and fetch the code from the repository:

```bash
mkdir $HOME/gocode
export GOPATH=$HOME/gocode
go get github.com/FineUploader/fineuploader-go-server
```

##### Build and install

```bash
cd $GOPATH/src/github.com/FineUploader/fineuploader-go-server
go install
```

##### Run the server

```bash
$GOPATH/bin/fineuploader-go-server
```

###### Open [**http://localhost:8080/**](http://localhost:8080/) in your browser

The default listening port is **8080** and the base upload directory is **uploads**, you can change that by passing the below optional flags to the executable,

```bash
$GOPATH/bin/fineuploader-go-server -p 9000 -d customUploadDir
```

### License ###
This project is licensed under the terms of the MIT license.
