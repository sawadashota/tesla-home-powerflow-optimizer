package middleware

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/contexts"
	"github.com/sawadashota/tesla-home-powerflow-optimizer/internal/logx"
)

type dumpResponseWriter struct {
	Code int
	http.ResponseWriter
}

func newDumpResponseWriter(w http.ResponseWriter) *dumpResponseWriter {
	return &dumpResponseWriter{
		ResponseWriter: w,
	}
}

func (w *dumpResponseWriter) WriteHeader(code int) {
	w.Code = code
	w.ResponseWriter.WriteHeader(code)
}

func (m *Middleware) NewLoggerMiddleware() func(http.Handler) http.Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logx.ParseLevel(m.r.AppConfig().LogLevel),
	}))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			started := time.Now().UTC()
			_, ctx := contexts.GetOrGenerateRequestID(r.Context())

			logger.Info("incoming request",
				slog.String("user-agent", r.UserAgent()),
				slog.String("time", started.UTC().Format(time.RFC3339)),
			)

			drw := newDumpResponseWriter(w)

			r = r.WithContext(ctx)
			next.ServeHTTP(drw, r)

			status := drw.Code
			if status == 0 {
				status = http.StatusOK
			}
			logger.Info("handled request",
				slog.Int("status", status),
				slog.String("user-agent", r.UserAgent()),
				slog.Int64("duration", time.Since(started).Nanoseconds()),
				slog.String("time", started.UTC().Format(time.RFC3339)),
			)
		})
	}
}
