package main

import (
	"fmt"
	"github.com/Kell9831/SpacedLearn/user"
)

type menuOption int 

const (
	UNSELECTED menuOption = iota
	REGISTER
	LOGIN
)

func main()  {

	fmt.Println("This is Spaced repetition app")
	fmt.Println("It helps you to learn english using the spaced repetition algorithm!!!")
	fmt.Println("Good luck!")

	var userAction menuOption

	for userAction == UNSELECTED {
		fmt.Println("I don't know who you are...")
		fmt.Println("1. Register")
		fmt.Println("2. Login")

		_,err := fmt.Scan(&userAction)
		if err != nil{
			userAction = UNSELECTED
		}

		switch userAction {
		case REGISTER:
			user.Register()
		case LOGIN:
			user.Login()

		default:
			userAction = UNSELECTED
			fmt.Println("Option not available")
		}
	}
}