package containers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"vivasoft-employee-entry-time-management/package/config"
	"vivasoft-employee-entry-time-management/package/connection"
	"vivasoft-employee-entry-time-management/package/controllers"
	"vivasoft-employee-entry-time-management/package/repositories"
	"vivasoft-employee-entry-time-management/package/routes"
	"vivasoft-employee-entry-time-management/package/services"
)

func Serve(echo *echo.Echo) {
	config.SetConfig()

	connection.Connect()
	db := connection.GetDB()

	// repository initialization
	employeeRepo := repositories.EmployeeDBInstance(db)

	// service initialization
	employeeService := services.EmployeeServiceInstance(employeeRepo)

	// controller initialization
	employeeCtr := controllers.NewEmployeeController(employeeService)

	// route initialization
	employeeRoutes := routes.EmployeeRoutes(echo, employeeCtr)

	employeeRoutes.InitEmployeeRoutes()

	// starting server
	log.Fatal(echo.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
