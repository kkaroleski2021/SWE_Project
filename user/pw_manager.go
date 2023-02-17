package user

import (
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
	//find password based on user's email...

	//empty user struct to store results of a query
	var resultUser = User{}
	DB.Where("email = ?", user.Email).Find(&resultUser)
	err := bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(user.Password))
	// nil == match
	if err != nil {
		panic(err.Error())
	}
	return err
}
