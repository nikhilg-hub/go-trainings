package main

import "fmt"

type Day int

// Declare an enum
const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// Implementation of stringer interface:
func (d Day) String() string {
	switch d {
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tueday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday, Sunday:
		return "Weekend"
	default:
		return "Invalid Weekday"
	}
}

func main() {

	D := Wednesday
	fmt.Println(D)

	W := Saturday

	fmt.Println(W)

}
