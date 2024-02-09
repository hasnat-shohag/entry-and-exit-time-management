package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vivasoft-employee-entry-time-management/package/domain"
	"vivasoft-employee-entry-time-management/package/types"
)

type EmployeeController struct {
	EmployeeSrv domain.IEmployeeService
}

func NewEmployeeController(EmployeeSrv domain.IEmployeeService) EmployeeController {
	return EmployeeController{
		EmployeeSrv: EmployeeSrv,
	}
}

func (emp *EmployeeController) GetAllEmployee(e echo.Context) error {
	employees, err := emp.EmployeeSrv.GetAllEmployee()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, employees)
}

func (emp *EmployeeController) GetEmployeeById(e echo.Context) error {
	tempEmpID := e.Param("id")
	employeeID, err := strconv.ParseInt(tempEmpID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Employee ID")
	}

	employee, err := emp.EmployeeSrv.GetEmployeeById(uint(employeeID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, employee)
}

func (emp *EmployeeController) CreateEmployee(e echo.Context) error {
	reqEmployee := &types.EmployeeRequest{}

	if err := e.Bind(reqEmployee); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}

	if err := reqEmployee.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := emp.EmployeeSrv.CreateEmployee(reqEmployee); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, "Employee created successfully")
}

func (emp *EmployeeController) DeleteEmployee(e echo.Context) error {
	id := e.Param("id")
	employeeID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid ID")
	}

	_, err = emp.EmployeeSrv.GetEmployeeById(uint(employeeID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := emp.EmployeeSrv.DeleteEmployee(uint(employeeID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "Employee deleted successfully")
}

func (emp *EmployeeController) UpdateEmployee(e echo.Context) error {

	tempEmpID := e.Param("id")
	empID, err := strconv.ParseInt(tempEmpID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Employee ID")
	}

	_, err = emp.EmployeeSrv.GetEmployeeById(uint(empID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	reqEmployee := &types.EmployeeRequest{}

	if err := e.Bind(reqEmployee); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}

	if err := reqEmployee.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	if err := emp.EmployeeSrv.UpdateEmployee(reqEmployee, uint(empID)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "Employee updated successfully")
}


