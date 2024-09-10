package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Kell9831/SpacedLearn/models"
	"github.com/Kell9831/SpacedLearn/utils"
)

func Register()  {
	var name string
	
	for{
		fmt.Println("Please tell me your name?")
		fmt.Scanln(&name)
		name = strings.TrimSpace(name)
		
		if len(name) < 4{
			fmt.Println("This is not a valid name. It must be at least 5 characters long.")
			continue
		}

		err := saveUser(models.User{Name: name})
		if err != nil {
			fmt.Println("Error registering the user:", err)
			continue
		}

		fmt.Println("Welcome to the Spaced Repetition App")
		break
	}
}

func saveUser(newUser models.User) error{
	users, err := utils.LoadUsers()

	if err != nil{
		return err
	}

	for _,user := range users {
		if strings.EqualFold(user.Name, newUser.Name){
			return errors.New("user already registered")
		}
	}

	users = append(users, newUser)

	return utils.CreateAndWriteToFile(users)
	
}


func Login()  {

	var name string

	users,err:= utils.LoadUsers()
	if err != nil{
		fmt.Println("Error loading users: ", err)
		return
	}

	for{
		fmt.Println("Please remember me your username?")
		fmt.Scanln(&name)
		name = strings.TrimSpace(name)

		for _,user := range users {
			if user.Name == name {
				fmt.Printf("Hello %s \n" , name)
				fmt.Println("Let's start learning")
				learn.Learn(name)
				break
			}
		}

		fmt.Println("Invalid user name")
	}
}