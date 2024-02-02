package services

import (
	"errors"
	"vivasoft-employee-entry-time-management/package/domain"
	"vivasoft-employee-entry-time-management/package/models"
)

type employeeService struct {
	employeeRepo domain.IEmployeeRepo
}

func EmployeeServiceInstance(employeeRepo domain.IEmployeeRepo) domain.IEmployeeService {
	return &employeeService{
		employeeRepo: employeeRepo,
	}
}

func (service *employeeService) GetAllEmployee() ([]models.Employee, error) {
	employee := service.employeeRepo.GetAllEmployee()
	if len(employee) == 0 {
		return nil, errors.New("employee not found")
	}
	return employee, nil
}

func (service *employeeService) GetEmployeeById(id uint) (models.Employee, error) {
	employeeDetail, err := service.employeeRepo.GetEmployeeById(id)

	if err != nil {
		return employeeDetail, errors.New("employee not found")
	}
	return employeeDetail, nil
}

func (service *employeeService) CreateEmployee(employee *models.Employee) error {
	if err := service.employeeRepo.CreateEmployee(employee); err != nil {
		return errors.New("employee not created")
	}
	return nil
}

func (service *employeeService) UpdateEmployee(updateEmployee *models.Employee) error {
	existingEmployee, err := service.GetEmployeeById(updateEmployee.ID)
	if err != nil {
		return errors.New("employee not found")
	}
	if updateEmployee.Name == "" {
		updateEmployee.Name = existingEmployee.Name
	}
	if updateEmployee.Designation == "" {
		updateEmployee.Designation = existingEmployee.Designation
	}
	if updateEmployee.Date.IsZero() {
		updateEmployee.Date = existingEmployee.Date
	}
	if updateEmployee.CreatedAt.IsZero() {
		updateEmployee.CreatedAt = existingEmployee.CreatedAt
	}
	if updateEmployee.EntryTime.IsZero() {
		updateEmployee.EntryTime = existingEmployee.EntryTime
	}
	if updateEmployee.ExitTime.IsZero() {
		updateEmployee.ExitTime = existingEmployee.ExitTime
	}

	if err := service.employeeRepo.UpdateEmployee(updateEmployee); err != nil {

		return errors.New("employee not updated")
	}

	return nil
}

func (service *employeeService) DeleteEmployee(id uint) error {
	if err := service.employeeRepo.DeleteEmployee(id); err != nil {
		return errors.New("employee is not deleted")
	}
	return nil
}