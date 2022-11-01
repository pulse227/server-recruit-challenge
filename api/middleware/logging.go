package middleware

import (
	"log"
	"net/http"
)

type loggingWriter struct {
	http.ResponseWriter
	code int
}

func newLoggingWriter(w http.ResponseWriter) *loggingWriter {
	return &loggingWriter{ResponseWriter: w, code: http.StatusInternalServerError}
}

func (lw *loggingWriter) WriteHeader(code int) {
	lw.code = code
	lw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("uri: %s, method: %s\n", req.RequestURI, req.Method)

		rlw := newLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Printf("response code: %d", rlw.code)
	})
}
