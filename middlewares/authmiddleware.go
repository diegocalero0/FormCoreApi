package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"formcoreapi/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("cae78716085a9f35684dc57f1e2c1e90104b666be03601e8cb10ecc420c2684b")

// Middleware that verifies if the authorization bearer token is valid
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			errorResponse := models.ErrorModel{
				Message: "Authentication Header Required",
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			errorResponse := models.ErrorModel{
				Message: "Invalid authorization header",
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {

			errorResponse := models.ErrorModel{
				Message: err.Error(),
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// This method creates a generic bearer token for process anothers request
func GenerateTokenHandler(w http.ResponseWriter, _ *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "Teamform generic username",
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"token":"%s"}`, tokenString)))
}
