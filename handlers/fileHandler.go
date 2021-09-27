package handlers

import (
	"FileWithAuth/models"
	"encoding/json"
	"net/http"
)

type fileHandler struct {
	filename string
	next     http.Handler
}

func (h *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dataMap := map[string]interface{}{}
	files := models.GetAllFileDataList()

	dataMap["FileList"] = files

	h.next = TemplateHandlerWithData(h.filename, dataMap)

	h.next.ServeHTTP(w, r)
}

func GetFileHandler(htmlfile string) http.Handler {

	return &fileHandler{filename: htmlfile}
}

func GetFileList(w http.ResponseWriter, req *http.Request) {
	files := models.GetAllFileDataList()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(files)
}
