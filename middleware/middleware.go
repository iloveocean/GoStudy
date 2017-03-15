package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

func loggingHandler(next http.Handler) http.Handler {
	fmt.Println("enter loggingHandler")
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("enter loggingHandler anonymous function")
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
		fmt.Println("leave loggingHandler amonymous function")
	}
	fmt.Println("leave loggingHandler")
	return http.HandlerFunc(fn)
}

func recoverHandler(next http.Handler) http.Handler {
	fmt.Println("enter recoverHandler")
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("enter recoverHandler anonymous function")
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
		fmt.Println("leave recoverHandler anonymous function")
	}
	fmt.Println("leave recoverHandler")
	return http.HandlerFunc(fn)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter about handler")
	fmt.Fprintf(w, "You are on the about page.")
	fmt.Println("leave about handler")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter index handler")
	fmt.Fprintf(w, "Welcome!")
	fmt.Println("leave index handler")
}

func main() {
	commonHandler := alice.New(loggingHandler, recoverHandler)
	http.Handle("/about", commonHandler.ThenFunc(aboutHandler))
	http.Handle("/", commonHandler.ThenFunc(indexHandler))
	http.ListenAndServe(":3000", nil)
}
