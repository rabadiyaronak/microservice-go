package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Opps Something went wrong!!!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s\n", data)
}
