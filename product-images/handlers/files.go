package handlers

import (
	"io"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/rabadiyaronak/microservice-go/product-images/files"
)

// Files is handler to read and write files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

//constructir

func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

//handle upload file via rest
func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	fileName := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fileName)

	f.saveFile(id, fileName, rw, r.Body)

}

//handle upload file with multipart request
func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Error("Bad Request", "error", err)
		http.Error(rw, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	id, idErr := strconv.Atoi(r.FormValue("id"))
	f.log.Info("process form for id:", "id", id)

	if idErr != nil {
		f.log.Error("Bad Request , form valud id requeired", "error", err)
		http.Error(rw, "Expected form vaue id", http.StatusBadRequest)
		return
	}

	ff, mh, err := r.FormFile("file")

	if err != nil {
		f.log.Error("Bad Request , file requeired", "error", err)
		http.Error(rw, "Expected file", http.StatusBadRequest)
		return
	}

	f.saveFile(r.FormValue("id"), mh.Filename, rw, ff)

}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)

	if err != nil {
		f.log.Error("Unable to store file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
