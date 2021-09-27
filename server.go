package main

import (
	"FileWithAuth/handlers"
	"FileWithAuth/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
)

func main() {
	err := godotenv.Load("env/auth.env")

	if err != nil {
		log.Fatal("Error loading rdsDB.env file", err)
	}

	router := mux.NewRouter().StrictSlash(true)
	googleClient := os.Getenv("GOOGLE_CLIENT_ID")
	googleSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New(googleClient, googleSecret,
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
