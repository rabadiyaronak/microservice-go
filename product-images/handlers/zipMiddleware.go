package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/hashicorp/go-hclog"
)

type GZipHandler struct {
	l hclog.Logger
}

func NewGZipHandler(logger hclog.Logger) *GZipHandler {
	return &GZipHandler{l: logger}
}

func (g *GZipHandler) GZipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			g.l.Info("gzip encoding supported.. continuing with compress data")
			wrw := NewWrappedResponseWriter(rw)
			wrw.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrw, r)

			defer wrw.Fulsh()

			return
		}

		//Handle request normally
		next.ServeHTTP(rw, r)
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gzipWriter := gzip.NewWriter(rw)

	return &WrappedResponseWriter{
		rw: rw,
		gw: gzipWriter,
	}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

func (wr *WrappedResponseWriter) Fulsh() {
	wr.gw.Flush()
	wr.gw.Close()
}
