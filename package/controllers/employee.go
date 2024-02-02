package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"vivasoft-employee-entry-time-management/package/domain"
	"vivasoft-employee-entry-time-management/package/models"
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
	books, err := emp.EmployeeSrv.GetAllEmployee()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, books)
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
	employee := &models.Employee{
		Name:        reqEmployee.Name,
		Designation: reqEmployee.Designation,
		Date:        reqEmployee.Date,
		EntryTime:   reqEmployee.EntryTime,
		ExitTime:    reqEmployee.ExitTime,
	}

	if err := emp.EmployeeSrv.CreateEmployee(employee); err != nil {
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
	reqEmployee := &types.EmployeeRequest{}

	if err := e.Bind(reqEmployee); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Data")
	}

	tempEmpID := e.Param("id")
	empID, err := strconv.ParseInt(tempEmpID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Employee ID")
	}

	_, err = emp.EmployeeSrv.GetEmployeeById(uint(empID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	updatedEmployee := &models.Employee{
		Name:        reqEmployee.Name,
		Designation: reqEmployee.Designation,
		Date:        reqEmployee.Date,
		EntryTime:   reqEmployee.EntryTime,
		ExitTime:    reqEmployee.ExitTime,
	}
	updatedEmployee.ID = uint(empID)

	if err := emp.EmployeeSrv.UpdateEmployee(updatedEmployee); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, "Employee updated successfully")
}