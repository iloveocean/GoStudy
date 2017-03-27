package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	JASON_API_IDENTIFIER = "application/vnd.api+json"
)

func WriteError(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", JASON_API_IDENTIFIER)
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(Errors{[]*Error{err}})
}

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("enter LoggingHandler\n")
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q, %v\n", r.Method, r.URL.String(), t2.Sub(t1))
		log.Printf("leave LoggingHandler\n")
	}
	return http.HandlerFunc(fn)
}

func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("enter RecoverHandler\n")
		defer func() {
			if err := recover(); err != nil {
				errMsg := fmt.Sprintf("%v", err)
				log.Printf("panic: %s\n", errMsg)
				WriteError(w, ErrInternalServer)
			}
		}()
		next.ServeHTTP(w, r)
		log.Printf("leave RecoverHandler\n")
	}
	return http.HandlerFunc(fn)
}
