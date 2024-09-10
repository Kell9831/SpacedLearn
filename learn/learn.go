package learn

import (
	"fmt"
	"math"
	"os"

	"github.com/Kell9831/SpacedLearn/models"
)

func Learn(nameUser string)  {
	
	defaultList := []models.Word{
		{Word: "learn" , Translation: "aprender", CountSuccesses: 0, RetentionRate: 0},
		{Word: "run" , Translation: "correr", CountSuccesses: 0, RetentionRate: 0},
		{Word: "sleep" , Translation: "dormir", CountSuccesses: 0, RetentionRate: 0},
		{Word: "eat" , Translation: "comer", CountSuccesses: 0, RetentionRate: 0},
		{Word: "walk" , Translation: "caminar", CountSuccesses: 0, RetentionRate: 0},
	}

	list := GetCurrentProgress(nameUser)

	if list == nil {
		list = defaultList
	}

	var answer string
	fmt.Printf("What does '%s' mean in Spanish?\n", list[0].Word)
	fmt.Scanln(&answer)

	if answer == list[0].Translation {
		newIndex := int(math.Pow(2, float64(list[0].CountSuccesses+1)))
		if newIndex >= len(list){
			newIndex = len(list) -1
		}

		orderedList := orderAsSpacedRepetition(list, newIndex)
		updatedList := UpdateRetentionRate(orderedList, newIndex)
		UpdateCurrentProgress(models.User{Name: nameUser, CurrentProgress: updatedList})
		printList(updatedList)
		fmt.Println("Congratulations! It's correct!")

	}else {
		fmt.Printf("Sorry, it's not correct.The correct for '%s' is : %s\n",list[0].Word ,list[0].Translation)
		list[0], list[1] = list[1], list[0]
		UpdateCurrentProgress(models.User{Name: nameUser, CurrentProgress: list})
		printList(list)
		
	}
	
	ShowOptions(nameUser)
}

func orderAsSpacedRepetition(list []models.Word , newIndex int) []models.Word  {
	
	tempList := make([]models.Word ,len(list))

	copy(tempList,list [1:newIndex + 1])
	tempList[newIndex] = list[0]
	copy(tempList[newIndex+1:], list[newIndex+1:])

	list = tempList
	list[newIndex].CountSuccesses++

	return list
}

func printList(list []models.Word) {
	fmt.Printf("::::: NEW LIST ::::: [")
	for i, word := range list {
		fmt.Printf("[%d]%s", i, word.Word)
		if i != len(list)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}

func ShowOptions(nameUser string) {
	fmt.Printf("What did you want to do?\n" +
		"1. Go to the next word\n" +
		"2. See my progress\n" +
		"3. Finish\n")

	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		Learn(nameUser)
	case 2:
		ShowProgress(nameUser)
	case 3:
		fmt.Printf("Thanks for practicing with the Space repetition app\n" +
			"See you later!\n")
		os.Exit(0)
	default:
		fmt.Printf("Invalid option\n")
		ShowOptions(nameUser)
	}
}
