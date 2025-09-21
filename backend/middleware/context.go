package middleware

import (
	"net/http"
)

type contextKey string

const UserIDKey = contextKey("userID")

func GetUserID(r *http.Request) string {
	userID, ok := r.Context().Value(UserIDKey).(string)
	if !ok {
		return ""
	}
	return userID
}
