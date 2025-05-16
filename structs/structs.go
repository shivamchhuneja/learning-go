package structs

import (
	"fmt"
	"time"

	"github.com/shivamchhuneja/learning-go/helpers"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	createdAt time.Time
}

func Struct_fn() {
	FirstName, _ := helpers.StrUserInput("Please enter your first name: ")
	LastName, _ := helpers.StrUserInput("Please enter your last name: ")
	BirthDate, _ := helpers.StrUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := User{
		FirstName: FirstName,
		LastName:  LastName,
		BirthDate: BirthDate,
		createdAt: time.Now(),
	}

	outputsUserDetails(&appUser)
}

func outputsUserDetails(u *User) {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}
