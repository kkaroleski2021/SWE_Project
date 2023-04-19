package user

// import (
// 	//"github.com/golang-jwt/jwt"
// 	"context"
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	//"os"
// 	//"strings"
// 	"time"
// 	//jwt "github.com/dgrijalva/jwt-go"
// 	//"github.com/gorilla/context"
// )

// type SessionInterface interface {
// 	isActive(cookie *http.Cookie) bool
// 	SetCookie(w http.ResponseWriter, r *http.Request, user *User)
// 	CheckCookie(w http.ResponseWriter, r *http.Request) int
// 	ValidateToken(next http.HandlerFunc) http.HandlerFunc
// }

// // we'll use this method later to determine if the session has expired
// func isActive(cookie *http.Cookie) bool {
// 	return cookie.Expires.Before(time.Now())
// }

// func SetCookie(w http.ResponseWriter, r *http.Request, user *User) {
// 	expiration := time.Now().Add(20 * time.Minute)
// 	cookie := http.Cookie{
// 		Name:    "logged-in",
// 		Value:   strconv.FormatUint(uint64(user.ID), 10),
// 		Expires: expiration,
// 		Path:    "/",
// 	}
// 	http.SetCookie(w, &cookie)
// }

// func CheckCookie(w http.ResponseWriter, r *http.Request) int {
// 	retVal := -1
// 	cookie, err := r.Cookie("logged-in")
// 	//if no cookie
// 	if err == http.ErrNoCookie {
// 		json.NewEncoder(w).Encode("No Cookie")
// 	} else {
// 		if isActive(cookie) {
// 			retVal, err = strconv.Atoi(cookie.Value)
// 		} else {
// 			json.NewEncoder(w).Encode("Session has expired")
// 		}
// 	}
// 	//returns the user id or -1 if not logged in / no session
// 	return retVal
// }

// func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Content-Type", "application/json")
// 		valid := CheckCookie(w, r)
// 		if valid == -1 {
// 			json.NewEncoder(w).Encode("No Valid Session")
// 			return
// 		}
// 		//write the id to the request to pass along to next function
// 		ctx := context.WithValue(r.Context(), "request_id", valid)
// 		next(w, r.WithContext((ctx)))
// 	})
// }
