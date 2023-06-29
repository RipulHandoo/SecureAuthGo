package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RipulHandoo/jwt/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.Jwtuser)

func Auth(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		godotenv.Load()
		var jwtKey string = os.Getenv("JWT_SECRET_KEY")
		jwtToken, err := req.Cookie("auth_token")

		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("No cookie")
				ResponseWithError(w, http.StatusUnauthorized, err)
				return
			}
			ResponseWithError(w, http.StatusUnauthorized, err)
			return
		}
		tknStr := jwtToken.Value
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			fmt.Println("No valid jwt")
			ResponseWithError(w, http.StatusUnauthorized, err)
			return

		if !tkn.Valid {
			ResponseWithError(w, http.StatusUnauthorized, err)
			return
		}
		userEmail := claims.Creds.Email
		apiConfig := DbClient

		user, dbErr2 := apiConfig.GetUserByEmail(req.Context(), userEmail)
		if dbErr2 != nil {
			ResponseWithError(w, http.StatusInternalServerError, dbErr2)
			return
		}

		handler(w, req, user)
	}
	}
}