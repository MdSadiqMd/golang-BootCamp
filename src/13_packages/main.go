package main

import (
	"fmt"
	"github.com/MdSadiqMd/golang-BootCamp/auth"
	"github.com/MdSadiqMd/golang-BootCamp/user"
)

func main() {
	auth.LoginWithCredentials("sadiq", "1234")

	userData := user.User{
		Name:  "Sadiq",
		Email: "sadiq@gmail",
	}
	fmt.Println(userData.Email, userData.Name)
}
