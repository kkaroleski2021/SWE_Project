package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Login interface {
	HashPassword(user *User)
	CheckPassword(user *User) error
}

func HashPassword(user *User) {
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(hashedPW)
}

func CheckPassword(user *User) error {
	//search for password in database
	//create a query for passwords in db
	var resultUser = User{Email: user.Email}
	DB.Select("password").Find(&resultUser)
	//find password based on user's email...
	
	err := bcrypt.CompareHashAndPassword( , []byte(user.Password))
	// nil == match
	if err != nil {
		panic(err.Error())
	}
	return err
}
