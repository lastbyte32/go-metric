package middleware

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	compressMethod = "gzip"
	compressLevel  = gzip.BestSpeed
)

func Compressor(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Content-Encoding") == compressMethod {
			reader, err := gzip.NewReader(r.Body)
			if err != nil {
				fmt.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer reader.Close()
			r.Body = reader
		}

		if !strings.Contains(r.Header.Get("Accept-Encoding"), compressMethod) {
			h.ServeHTTP(w, r)
			return
		}

		writer, err := gzip.NewWriterLevel(w, compressLevel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer writer.Close()

		w.Header().Set("Content-Encoding", compressMethod)

		h.ServeHTTP(compressWriter{
			ResponseWriter: w,
			Writer:         writer,
		}, r,
		)
	})
}

type compressWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w compressWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
