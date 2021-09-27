package main

import (
	"FileWithAuth/handlers"
	"FileWithAuth/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New("877820034342-se32uu9p26isr66946j3onarm05auisn.apps.googleusercontent.com", "RsMUdt1W5ORNmt6RQlhLpVfK",
			"http://localhost:8080/auth/callback/google"),
	)

	router.Handle("/login", handlers.TemplateHandler("login.html"))
	router.Handle("/upload", handlers.MustAuth(handlers.TemplateHandler("upload.html")))
	router.Handle("/", handlers.MustAuth(handlers.GetFileHandler("fileList.html")))
	router.Handle("/log/{filename}", handlers.MustAuth(handlers.GetFileLogHandler("fileLogList.html")))

	router.Handle("/uploader", handlers.MustAuth(handlers.Uploader()))
	router.Handle("/download/{filename}", handlers.Downloader())
	router.HandleFunc("/logout", handlers.LogoutHandler)
	router.HandleFunc("/auth/{action}/{provider}", handlers.LoginHandler)
	router.HandleFunc("/delete", func(w http.ResponseWriter, req *http.Request) { models.DeleteAll() })
	router.HandleFunc("/file", handlers.GetFileList)
	router.HandleFunc("/log", handlers.GetAllLogList)
	//router.HandleFunc("/log", func(w http.ResponseWriter, req *http.Request) { models.GetAllLogDataList() })
	log.Fatal(http.ListenAndServe(":8080", router))
}
