package utils

import (
	"encoding/json"
	"os"

	"github.com/Kell9831/SpacedLearn/models"
)

var UserFile = "users.json"

func LoadUsers() ([]models.User,error)  {
	users := []models.User{} //Inicialización de la lista de usuarios

	file, err := os.Open(UserFile)
	if err != nil{
		if os.IsNotExist(err){
			return users, nil
		}
		return nil,err
	}

	err = json.NewDecoder(file).Decode(&users)
	if err != nil{
		return nil,err
	}

	return users, nil
}


func CreateAndWriteToFile(users []models.User)  error{
	file, err := os.Create(UserFile)
	if err != nil{
		return err
	}

	defer file.Close()

	data, err := json.Marshal(users)
	if err !=nil{
		return err
	}

	_,err = file.Write(data)
	return err
}