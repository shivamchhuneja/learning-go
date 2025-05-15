package structs

import (
	"fmt"
	"time"

	"github.com/shivamchhuneja/learning-go/helpers"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func Struct_fn() {
	userfirstName, _ := helpers.StrUserInput("Please enter your first name: ")
	userlastName, _ := helpers.StrUserInput("Please enter your last name: ")
	userBirthdate, _ := helpers.StrUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := User{
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	outputsUserDetails(appUser)
}

func outputsUserDetails(u User) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}
