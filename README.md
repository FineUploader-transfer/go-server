# Go (a.k.a. Golang) Server-Side Example for the Widen Fine Uploader Javascript Library #

This repository contains a [**Golang**](https://golang.org/) server-side example for users of [**Widen's Fine Uploader javascript library**](http://fineuploader.com/).  

### This server supports

* File chunking
* Concurrent uploading
* Delete uploaded files
* Pause / Resume uploads

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
POST|/upload|Upload file end point. Will create `<root_upload_directory>/qquuid` directory and store the received file inside
DELETE|/upload|Deletes the uploaded file based on the `qquuid` parameter
POST|/chunksdone|Builds original file based on received chunks for the received `qquuid` parameter


### FineUploader configuration

#### Download fineuploader

You can use `npm` to download fineuploader using the provided `package.json` file

```bash
$ cd $GOPATH/src/github.com/FineUploader/fineuploader-go-server
$ npm install
```

#### Configure upload endpoint

```javascript
request: {
    endpoint: 'upload'
}
```

#### Configure file chunking (optional)

```javascript
chunking: {
	enabled: true,
	concurrent: {
	    enabled: true
	},
	success: {
	    endpoint: 'chunksdone'
	}
}
```

#### Configure file deletes (optional)

```javascript
deleteFile: {
	enabled: true,
	endpoint: 'upload'
}
```

#### Enable ability to resume (optional)

```javascript
resume: {
    enabled: true
}
```

### License ###
This project is licensed under the terms of the MIT license.
