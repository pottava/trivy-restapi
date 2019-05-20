package lib

import (
	"log"
	"net/http"
	"time"
)

// Wrap wraps HTTP request handler
func Wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proc := time.Now()
		handler.ServeHTTP(w, r)
		if Config.AccessLog {
			log.Printf("[%s] %.3f %s %s", r.RemoteAddr, time.Since(proc).Seconds(), r.Method, r.URL)
		}
	})
}
