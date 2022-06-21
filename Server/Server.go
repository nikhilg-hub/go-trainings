package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Models
type Employee struct {
	Id   int
	Name string
	Age  int
}

type ResponseWithMessage struct {
	Message string
}

// In memory DB:
var EMP_DB = map[int]Employee{
	1: Employee{1, "Nishant", 23},
	2: Employee{2, "Sanchita", 24},
	3: Employee{3, "Shubham", 25},
}

// Auto increment
var AutoInc = 3

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	e.GET("/employee/:id", getEmployee)

	e.GET("/employee", getAllEmployees)

	e.POST("/employee", saveEmployee)

	e.DELETE("/employee/:id", deleteEmployee)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Sample Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "This is the REST api server!")
}

// List all employees
func getAllEmployees(c echo.Context) error {

	// Retrieve query parameter
	age := c.QueryParam("age")

	if age != "" {
		// convert age string to int
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			// handle conversion error
			return c.JSON(http.StatusBadRequest, ResponseWithMessage{"Cannot convert " + age + " to integer."})
		}

		empList := []Employee{}
		for _, v := range EMP_DB {
			// Filter employees with supplied age
			if v.Age == ageInt {
				empList = append(empList, v)
			}
		}
		return c.JSON(http.StatusOK, empList)
	}

	// return a full list of employees
	empList := []Employee{}

	for _, v := range EMP_DB {
		empList = append(empList, v)
	}

	return c.JSON(http.StatusOK, empList)
}

func getEmployee(c echo.Context) error {

	//read the path parameter:
	id := c.Param("id")

	// c.Logger().Println("id ------> ", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseWithMessage{"Cannot convert " + id + " to integer."})
	}

	if emp, exists := EMP_DB[idInt]; exists {
		return c.JSON(http.StatusOK, emp)
	}

	return c.NoContent(http.StatusNoContent)
}

func saveEmployee(c echo.Context) error {
	emp := new(Employee)

	if err := c.Bind(emp); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseWithMessage{"Error parsing the request body."})
	}

	// save logic:
	AutoInc++
	emp.Id = AutoInc
	EMP_DB[AutoInc] = *emp

	return c.JSON(http.StatusCreated, emp)

}

func deleteEmployee(c echo.Context) error {
	id := c.Param("id")

	// c.Logger().Println("id ------> ", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseWithMessage{"Cannot convert " + id + " to integer."})
	}

	if emp, exists := EMP_DB[idInt]; exists {
		delete(EMP_DB, idInt)
		return c.JSON(http.StatusOK, emp)
	}

	return c.NoContent(http.StatusNoContent)
}
