package handlers

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/objx"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
	data     map[string]interface{}
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("views", t.filename)))
	})

	t.data["Host"] = r.Host

	if authCookie, err := r.Cookie("auth"); err == nil {
		t.data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, t.data)
}

func TemplateHandlerWithData(file string, data map[string]interface{}) http.Handler {
	templateHandle := &templateHandler{filename: file, data: map[string]interface{}{}}

	for key, value := range data {
		templateHandle.data[key] = value
	}

	return templateHandle
}

func TemplateHandler(file string) http.Handler {
	return &templateHandler{filename: file, data: map[string]interface{}{}}
}
