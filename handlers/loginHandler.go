package handlers

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	action := param["action"]
	provider := param["provider"]

	switch action {
	case "login":
		//Login Action Handler
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get provider %s: %s", provider, err), http.StatusBadRequest)
			return
		}
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to GetBeginAuthURL for %s:%s", provider, err), http.StatusInternalServerError)
			return
		}
		log.Println(loginUrl)
		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)
	case "callback":
		//Callback Action Handler
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to complete auth for %s : %s", provider, err), http.StatusBadRequest)
			return
		}
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get user from %s : %s", provider, err), http.StatusInternalServerError)
			return
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error when trying to get user from %s : %s", provider, err), http.StatusInternalServerError)
			return
		}

		m := md5.New()
		io.WriteString(m, strings.ToLower(user.Email()))
		uniqueID := fmt.Sprintf("%x", m.Sum(nil))

		if err != nil {
			log.Fatalln("Error when trying to GetAvatarURL", "-", err)
		}

		authCookieValue := objx.New(map[string]interface{}{
			"userid": uniqueID,
			"name":   user.Name(),
		}).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	w.Header().Set("Location", "/login")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
