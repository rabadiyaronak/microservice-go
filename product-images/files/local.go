package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

//local is implementation of Storage interface
//Use local disk on current machine
type Local struct {
	maxFileSize int //max number of bytes for a file
	basepath    string
}

//constructor for Local
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)

	if err != nil {
		return nil, err
	}

	return &Local{basepath: p, maxFileSize: maxSize}, nil
}

func (l *Local) Save(path string, content io.Reader) error {
	//Get full path of file
	fp := l.fullPath(path)

	dir := filepath.Dir(fp)

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		return xerrors.Errorf("Unable to create directory %w", err)
	}

	//check if file already exists
	_, err = os.Stat(fp)

	if err == nil {
		//delete already exists file
		err = os.Remove(fp)

		if err != nil {
			return xerrors.Errorf("Unable to delete file %w", err)
		}

		//if there is an error other than file already exists
	} else if !os.IsNotExist(err) {
		return xerrors.Errorf("Unable to get fileInfo %w", err)
	}

	//create file
	f, err := os.Create(fp)
	if err != nil {
		return xerrors.Errorf("Unable to create file %w", err)
	}

	defer f.Close()

	//write contents to file
	_, err = io.Copy(f, content)
	if err != nil {
		return xerrors.Errorf("Unable to write to file %w", err)
	}

	return nil

}

// Get the file at the given path and return a Reader
// the calling function is responsible for closing the reader
func (l *Local) Get(path string) (*os.File, error) {
	// get the full path for the file
	fp := l.fullPath(path)

	// open the file
	f, err := os.Open(fp)
	if err != nil {
		return nil, xerrors.Errorf("Unable to open file: %w", err)
	}

	return f, nil
}

func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basepath, path)
}
