package middleware

import (
	"log/slog"
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
		slog.InfoContext(req.Context(), "access", "uri", req.RequestURI, "method", req.Method)

		rlw := newLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		slog.InfoContext(req.Context(), "response", "code", rlw.code)
	})
}
