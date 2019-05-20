package lib

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Wrap wraps HTTP request handler
func Wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case eqauls(r, "/health"):
			w.WriteHeader(http.StatusOK)

		case eqauls(r, "/version"):
			fmt.Fprintf(w, "%s", Config.Version)

		default:
			proc := time.Now()
			handler.ServeHTTP(w, r)
			if Config.AccessLog {
				log.Printf("[%s] %.3f %s %s", r.RemoteAddr, time.Since(proc).Seconds(), r.Method, r.URL)
			}
		}
	})
}

func eqauls(r *http.Request, url string) bool {
	return url == r.URL.Path
}
