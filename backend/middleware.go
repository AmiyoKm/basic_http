package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
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

type contextKey string

func authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "no auth header found", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "malformed auth header", http.StatusUnauthorized)
			return
		}

		accessToken := parts[1]

		jwtPayload, err := JWTVerify(accessToken, SECRET)
		if err != nil {
			http.Error(w, "invalid jwt payload", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextKey("userID"), jwtPayload.Sub)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
