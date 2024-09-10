package learn

import (
	"fmt"

	"github.com/Kell9831/SpacedLearn/models"
	"github.com/Kell9831/SpacedLearn/utils"
)

func GetCurrentProgress(nameUser string) []models.Word  {
	users , err := utils.LoadUsers()

	if err != nil{
		fmt.Println("Error loading users:" , err)
		return []models.Word{} //devuelve un slice vacio si hay error
	}

	for _,user := range users {
		if user.Name == nameUser{
			return user.CurrentProgress //obtiene el progreso actual
		}
	}

	return []models.Word{}
}

func UpdateRetentionRate(list []models.Word, newIndex int) []models.Word {
	if list[newIndex].CountSuccesses <= 100 {
		list[newIndex].RetentionRate += 25
	}
	return list
}

func ShowProgress(nameUser string) {
	list := GetCurrentProgress(nameUser)
	if list == nil {
		fmt.Println("No progress to show")
		return
	}

	fmt.Println("Your current progress is:")
	fmt.Println("Words        | Retention Rate")
	for _, word := range list {
		fmt.Printf("%-12s | %3d%%\n", word.Word, word.RetentionRate)
	}

	ShowOptions(nameUser)
}

func UpdateCurrentProgress(userToEdit models.User) error {
	users, err := utils.LoadUsers()
	if err != nil {
		return err
	}

	// Encontrar el usuario que modificaremos y actualizarlo
	for i, user := range users {
		if user.Name == userToEdit.Name {
			users[i] = userToEdit
			break
		}
	}

	// Escribir las actualizaciones
	err = utils.CreateAndWriteToFile(users)
	return err
}
