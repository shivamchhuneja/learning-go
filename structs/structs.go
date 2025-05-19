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

type Admin struct {
	Email    string
	Password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func NewUser(FirstName, LastName, BirthDate string) User {
	return User{
		FirstName: FirstName,
		LastName:  LastName,
		BirthDate: BirthDate,
		createdAt: time.Now(),
	}
}

func Struct_fn() {
	UserFirstName, _ := helpers.StrUserInput("Please enter your first name: ")
	UserLastName, _ := helpers.StrUserInput("Please enter your last name: ")
	UserBirthDate, _ := helpers.StrUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := NewUser(UserFirstName, UserLastName, UserBirthDate)

	outputsUserDetails(&appUser)

	adminst := NewAdmin("test@test.com", "abc@123")

	outputsUserDetails(&adminst.User)
}

func outputsUserDetails(u *User) {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}
