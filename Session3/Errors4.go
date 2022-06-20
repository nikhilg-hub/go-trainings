package main

import (
	"errors"
	"fmt"
)

// Router
func main() {
	emp, err := FindEmployeeById(10)

	if err != nil {
		// fmt.Println(errors.Unwrap(err))
		var dbError DBError
		switch {
		case errors.As(err, &dbError):
			fmt.Println("error DB:", dbError)
		default:
			fmt.Println("some other error: ", err.Error())
		}
		return
	}

	fmt.Println(emp)

}

// Service
func FindEmployeeById(id int) (Employee, error) {
	employee, err := GetEmployeeFromDB(id)
	// error case
	if err != nil {
		return employee, fmt.Errorf("cannot find employee with error: %w", err)
	}
	// success case
	return Employee{10, "ABC"}, nil
}

// DAL
func GetEmployeeFromDB(id int) (Employee, error) {
	return Employee{}, DBError{}
}

type Employee struct {
	Id   int
	Name string
}

// DB layer
type DBError struct{}

func (DBError) Error() string {
	return "Error accessing the database"
}

// // Service layer
// type EmployeeNotFoundError struct {
// 	Id int
// 	// BaseError error // DB Error
// }

// func (err *EmployeeNotFoundError) Error() string {
// 	return fmt.Sprintf("cannot find employee with id : %d", err.Id)
// }
