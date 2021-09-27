package handlers

import (
	"FileWithAuth/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type fileLogHandler struct {
	filename string
	next     http.Handler
}

func (h *fileLogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filename := params["filename"]
	dataMap := map[string]interface{}{}
	logs := models.GetLogDataList(filename)

	dataMap["LogList"] = logs
	h.next = TemplateHandlerWithData(h.filename, dataMap)

	h.next.ServeHTTP(w, r)
}

func GetFileLogHandler(htmlfile string) http.Handler {
	return &fileLogHandler{filename: htmlfile}
}

func GetLogList(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	filename := params["filename"]
	logs := models.GetLogDataList(filename)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logs)
}

func GetAllLogList(w http.ResponseWriter, req *http.Request) {
	logs := models.GetAllLogDataList()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logs)
}
