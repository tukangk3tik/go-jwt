package middleware

import (
	"go-jwt/helper"
	"go-jwt/service"
	"net/http"
	"strings"
)

func AuthorizeJwt(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.ReturnFailResponse(w, "Please put your token", http.StatusUnauthorized)
			return
		}

		strToken := strings.Split(authHeader, " ")
		validateRes, err := service.JwtServiceInstance.ValidateToken(strToken[1])
		if err != nil {
			helper.ReturnFailResponse(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		r.Header.Set("email", validateRes.(string))
		f(w, r)
	}
}
