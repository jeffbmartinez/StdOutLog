package stdoutlog

import (
	"log"
	"net/http"
	"time"
)

type Middleware struct{}

func (m Middleware) ServeHTTP(response http.ResponseWriter, request *http.Request, next http.HandlerFunc) {
	start := time.Now()

	log.Printf("%v - %v (from %v)\n", request.Method, request.URL.String(), request.RemoteAddr)
	for key, value := range request.Header {
		log.Printf("%v: %v\n", key, value)
	}

	next(response, request)

	log.Printf("Completed in %v\n\n", time.Since(start))
}
