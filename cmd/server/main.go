package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Vladeeg/gometrics/internal/handler"
)

type Middleware func(http.Handler) http.Handler

func Conveyor(h http.Handler, middlewares ...Middleware) http.Handler {
    for _, middleware := range middlewares {
        h = middleware(h)
    }
    return h
} 

func HandleBadRequest(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusBadRequest)
}

func HandleNotFound(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
}

func isPostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST method allowed", http.StatusMethodNotAllowed)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func hasNameMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")

		if name == "" {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	portPtr := flag.Int("p", 8080, "server port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/update/gauge/{name}/{value}", Conveyor(http.HandlerFunc(handler.HandleGauge), isPostMiddleware, hasNameMiddleware))
	mux.Handle("/update/counter/{name}/{value}", Conveyor(http.HandlerFunc(handler.HandleCount), isPostMiddleware, hasNameMiddleware))
	mux.HandleFunc("/update/{type}/{$}", HandleNotFound)
	mux.HandleFunc("/update/", HandleBadRequest)
	mux.HandleFunc("/", HandleNotFound)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), mux))
}
