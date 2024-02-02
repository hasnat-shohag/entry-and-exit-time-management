package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"vivasoft-employee-entry-time-management/package/controllers"
)

type employeeRoutes struct {
	echo        *echo.Echo
	employeeCtr controllers.EmployeeController
}

func EmployeeRoutes(echo *echo.Echo, employeeCtr controllers.EmployeeController) *employeeRoutes {
	return &employeeRoutes{
		echo:        echo,
		employeeCtr: employeeCtr,
	}
}

func (employee *employeeRoutes) InitEmployeeRoutes() {
	e := employee.echo
	employee.initEmployeeRoutes(e)
}

func (emp *employeeRoutes) initEmployeeRoutes(e *echo.Echo) {
	//grouping route endpoints
	employee := e.Group("/employees")

	employee.GET("/ping", Pong)

	//initializing http methods - routing endpoints and their handlers
	employee.GET("/get-all-employees", emp.employeeCtr.GetAllEmployee)
	employee.GET("/get-employee/:id", emp.employeeCtr.GetEmployeeById)
	employee.POST("/create-employee", emp.employeeCtr.CreateEmployee)
	employee.PUT("/update-employee/:id", emp.employeeCtr.UpdateEmployee)
	employee.DELETE("/delete-employee/:id", emp.employeeCtr.DeleteEmployee)
}

func Pong(ctx echo.Context) error {
	fmt.Println("Pong")
	return ctx.JSON(http.StatusOK, "Pong")
}
