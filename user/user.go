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
		name = strings.ToLower(name)
		
		if len(name) < 5{
			fmt.Println("This is not a valid name. It must be at least 5 characters long.")
			continue
		}

		

	}
}

func saveUser(newUser models.User) error{
	users, err := utils.LoadUsers()

	if err != nil{
		return err
	}

	for _,user := range users {
		if strings.ToLower(user.Name) == strings.ToLower(newUser.Name){
			return errors.New("user already registered")
		}
	}

	users = append(users, newUser)

	return utils.CreateAndWriteToFile(users)
	
}


func Login()  {
	
}