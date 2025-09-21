package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AmiyoKm/basic_http/jwt"
	"github.com/AmiyoKm/basic_http/utils"
)

type Middleware func(http.Handler) http.Handler

func (m *Manager) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.CORSHeader(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (m *Manager) Logger(next http.Handler) http.Handler {
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

func (m *Manager) Authentication(next http.Handler) http.Handler {
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

		jwtPayload, err := jwt.JWTVerify(accessToken, m.cfg.JWTSecretKey)
		if err != nil {
			http.Error(w, "invalid jwt payload", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, jwtPayload.Sub)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
