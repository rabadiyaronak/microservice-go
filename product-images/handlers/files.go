package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/rabadiyaronak/microserive-go/product-images/files"
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

func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	fileName := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fileName)

	f.saveFile(id, fileName, rw, r)

}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)

	if err != nil {
		f.log.Error("Unable to store file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
