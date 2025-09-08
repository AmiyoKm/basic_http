package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CORSHeader(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		curr := time.Now()
		method := r.Method
		route := r.URL.Path

		next.ServeHTTP(w, r)

		diff := time.Since(curr)

		content := fmt.Sprintf("TIME : %s , METHOD : %s , ROUTE : %s , DURATION : %s", curr.Format(time.RFC1123), method, route, diff.String())
		log.Print(content)
	})
}

func mid1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("i am middleware 1 that was supposed to run first")
		next.ServeHTTP(w, r)
		log.Println("i am middleware 1 that was supposed to run first , while returning")
	})
}

func mid2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("i am middleware 2 that was supposed to run first")
		next.ServeHTTP(w, r)
		log.Println("i am middleware 2 that was supposed to run first , while returning")
	})
}
