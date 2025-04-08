package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Ethea2/Distributed-Fault-Tolerance/services/common/pkg/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, h *http.Request) {
		reqToken := h.Header.Get("Authorization")

		if reqToken == "" {
			http.Error(w, fmt.Sprintf("%v. %v", http.StatusText(403), "No token was found."), 403)
			return
		}

		split := strings.Split(reqToken, "Bearer ")

		if len(split) == 0 {
			http.Error(w, fmt.Sprintf("%v. %v", http.StatusText(401), "No token was found."), 401)
			return
		}

		tokenstring := split[1]

		user := auth.DecodeToken(tokenstring)

		ctx := context.WithValue(h.Context(), "custom_claims", user)

		next.ServeHTTP(w, h.WithContext(ctx))
	})
}
