package domain

import (
	"vivasoft-employee-entry-time-management/package/models"
	"vivasoft-employee-entry-time-management/package/types"
)

type IEmployeeRepo interface {
	GetAllEmployee() []models.Employee
	GetEmployeeById(id uint) (models.Employee, error)
	CreateEmployee(employee *models.Employee) error
	UpdateEmployee(employee *models.Employee) error
	DeleteEmployee(id uint) error
}

type IEmployeeService interface {
	GetAllEmployee() ([]models.Employee, error)
	GetEmployeeById(id uint) (models.Employee, error)
	CreateEmployee(employee *types.EmployeeRequest) error
	DeleteEmployee(id uint) error
	UpdateEmployee(employee *types.EmployeeRequest, id uint) error
}
