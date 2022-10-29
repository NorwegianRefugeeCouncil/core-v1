package middleware

import (
	"bufio"
	"net"
	"net/http"
	"time"

	"github.com/nrc-no/notcore/internal/logging"
	"go.uber.org/zap"
)

func RequestLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		l := logging.NewLogger(r.Context())
		url := *r.URL

		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			host = r.RemoteAddr
		}

		if r.Header.Get("X-Forwarded-For") != "" {
			host = r.Header.Get("X-Forwarded-For")
		}

		start := time.Now().UTC()
		path := r.URL.Path
		if path == "" {
			path = "/"
		}

		uri := r.RequestURI
		if r.ProtoMajor == 2 && r.Method == "CONNECT" {
			uri = r.Host
		}
		if uri == "" {
			uri = url.RequestURI()
		}

		rw := &responseWriter{ResponseWriter: w}

		fields := []zap.Field{
			zap.String("ip", host),
			zap.String("method", r.Method),
			zap.Time("ts", start),
			zap.String("uri", uri),
			zap.String("proto", r.Proto),
			zap.String("user_agent", r.UserAgent()),
		}
		if l.Core().Enabled(zap.DebugLevel) {
			for k, h := range r.Header {
				fields = append(fields, zap.String("header."+k, h[0]))
			}
		}

		requestLogger := l.With(fields...)
		requestLogger.Info("request started")

		h.ServeHTTP(rw, r)

		postFields := []zap.Field{
			zap.Int("status", rw.statusCode),
			zap.Int64("duration", time.Since(start).Milliseconds()),
		}

		if l.Core().Enabled(zap.DebugLevel) {
			for k, h := range rw.header {
				postFields = append(postFields, zap.String("header."+k, h[0]))
			}
		}

		responseLogger := l.With(postFields...)
		responseLogger.Info("request completed")
	})
}

type responseWriter struct {
	http.ResponseWriter
	header     http.Header
	statusCode int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriter) Header() http.Header {
	header := w.ResponseWriter.Header()
	w.header = header
	return header
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
