package middleware

import (
	"net/http"
	"task_manager_api/internal/logger"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		logger.Log.WithFields(map[string]interface{}{
			"method":   r.Method,
			"path":     r.RequestURI,
			"duration": duration.String(),
			"remote":   r.RemoteAddr,
			"agent":    r.UserAgent(),
		}).Info("Handled request")
	})
}
