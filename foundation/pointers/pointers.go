package pointer_learn

import "fmt"

func pointers() {
	age := 32

	agePointer := &age

	fmt.Println(*agePointer) // dereferncing the pointer

	// adultYears := getAdultYears(age)
	// fmt.Println(adultYears)
}

func getAdultYears(age *int) int {
	return *age - 18
}
