package handlers

import (
	"FileWithAuth/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/objx"
)

type DownloadHandler struct {
}

func (h *DownloadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	param := mux.Vars(req)
	filename := param["filename"]
	authCookieVal := map[string]interface{}{}

	if authCookie, err := req.Cookie("auth"); err == nil {
		authCookieVal = objx.MustFromBase64(authCookie.Value)
	} else {
		log.Fatal("Error when retrieving cookie", err)
	}

	file := models.GetFileData(filename)

	models.InsertLogData(models.LogData{FileName: filename, User: authCookieVal["name"].(string), Action: "Download", Time: time.Now()})

	w.Header().Set("Location", file.FileURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func Downloader() http.Handler {
	return &DownloadHandler{}
}
