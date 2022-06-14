package main

import "fmt"

// Unexported:
var i = 44
var str string = "string"
var (
	name           = str
	age            = 76
	salary float64 = 67
)

// Exported:
const Pi = 3.14
const Greeting = "hello"

const (
	RateOfInterest     = 11.9
	DefaultYears   int = 10
)

func main() {

	shortHand := "shortHand"

	var varWithDataType string = "varWithDataType"

	var withoutDataType = "withoutDataType"

	fmt.Println(shortHand, varWithDataType, withoutDataType)

	fmt.Println(RateOfInterest)
}
