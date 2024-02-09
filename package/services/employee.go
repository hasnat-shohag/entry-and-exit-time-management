package services

import (
	"errors"
	"vivasoft-employee-entry-time-management/package/domain"
	"vivasoft-employee-entry-time-management/package/models"
	"vivasoft-employee-entry-time-management/package/types"
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

func (service *employeeService) CreateEmployee(reqEmployee *types.EmployeeRequest) error {

	employee := &models.Employee{
		Name:        reqEmployee.Name,
		Designation: reqEmployee.Designation,
		Date:        reqEmployee.Date,
		EntryTime:   reqEmployee.EntryTime,
		ExitTime:    reqEmployee.ExitTime,
	}

	if err := service.employeeRepo.CreateEmployee(employee); err != nil {
		return errors.New("employee not created")
	}
	return nil
}

func (service *employeeService) UpdateEmployee(reqEmployee *types.EmployeeRequest, ID uint) error {

	updatedEmployee := &models.Employee{
		Name:        reqEmployee.Name,
		Designation: reqEmployee.Designation,
		Date:        reqEmployee.Date,
		EntryTime:   reqEmployee.EntryTime,
		ExitTime:    reqEmployee.ExitTime,
	}
	updatedEmployee.ID = ID
	existingEmployee, err := service.GetEmployeeById(ID)
	if err != nil {
		return errors.New("employee not found")
	}

	if updatedEmployee.Name == "" {
		updatedEmployee.Name = existingEmployee.Name
	}
	if updatedEmployee.Designation == "" {
		updatedEmployee.Designation = existingEmployee.Designation
	}
	if updatedEmployee.Date.IsZero() {
		updatedEmployee.Date = existingEmployee.Date
	}
	if updatedEmployee.CreatedAt.IsZero() {
		updatedEmployee.CreatedAt = existingEmployee.CreatedAt
	}
	if updatedEmployee.EntryTime.IsZero() {
		updatedEmployee.EntryTime = existingEmployee.EntryTime
	}
	if updatedEmployee.ExitTime.IsZero() {
		updatedEmployee.ExitTime = existingEmployee.ExitTime
	}

	if err := service.employeeRepo.UpdateEmployee(updatedEmployee); err != nil {
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
