package user

import (
	//"github.com/golang-jwt/jwt"
	//"encoding/json"
	"fmt"
	"net/http"

	//"os"
	//"strings"
	"time"
	//jwt "github.com/dgrijalva/jwt-go"
	//"github.com/gorilla/context"
)

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

type session struct {
	username string
	expiry   time.Time
}

// we'll use this method later to determine if the session has expired
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

func SetCookie(w http.ResponseWriter, user *User) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: user.Email, Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Fprint(w, cookie)
}
