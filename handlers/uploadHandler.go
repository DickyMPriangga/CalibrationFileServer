package handlers

import (
	"FileWithAuth/S3"
	"FileWithAuth/models"
	"fmt"
	"log"
	"net/http"
	"time"
)

type UploadHandler struct {
}

func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	expire := fmt.Sprintf("%s 00:00", req.FormValue("expire"))
	timeExpire, _ := time.Parse("2006-01-02 15:04", expire)
	file, header, err := req.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status, filepath := S3.UploadFile(file, header)

	if !status {
		log.Fatal("Failed to upload file")
	}

	models.InsertFileData(models.FileData{FileName: header.Filename, FileAuthor: name, FileURL: filepath, Expire: timeExpire})
	models.InsertLogData(models.LogData{FileName: header.Filename, User: name, Action: "Upload", Time: time.Now()})

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func Uploader() http.Handler {
	return &UploadHandler{}
}
