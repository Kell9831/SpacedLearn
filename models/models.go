package models

type User struct{
	Name string
	CurrentProgress []Word 
}

type Word struct {
	Word string
	Translation string
	CountSuccesses int
	RetentionRate int
}
