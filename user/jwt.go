package user

import (
	//"github.com/golang-jwt/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

type JwtToken struct {
	AccessToken string `json:"access-token"`
}

var jwt_secret = os.Getenv("jwt_secret")

func CreateToken(w http.ResponseWriter, user *User) {
	// access token ttl
	ttl := 10 * time.Minute
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   ttl,
	})

	tokenString, error := token.SignedString([]byte(jwt_secret))
	if error != nil {
		fmt.Println(error)
	}
	response := map[string]interface{}{"status": true,
		"message": "success"}
	response["data"] = JwtToken{AccessToken: tokenString}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	return
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		authHeader := req.Header.Get("authorization")
		if authHeader != "" {
			bearerToken := strings.Split(authHeader, " ")
			token, error := jwt.Parse(bearerToken[0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(jwt_secret), nil
			})
			if error != nil {
				json.NewEncoder(w).Encode(error.Error())
				return
			}
			if token.Valid {
				context.Set(req, "decoded", token.Claims)
				next(w, req)
			} else {
				json.NewEncoder(w).Encode("Invalid Authorization Token")
				return
			}
		} else {
			json.NewEncoder(w).Encode("An authorization header is required")
			return
		}
	})
}
