package repositories

import (
	"gorm.io/gorm"
	"vivasoft-employee-entry-time-management/package/domain"
	"vivasoft-employee-entry-time-management/package/models"
)

type EmployeeRepo struct {
	db *gorm.DB
}

func EmployeeDBInstance(d *gorm.DB) domain.IEmployeeRepo {
	return &EmployeeRepo{
		db: d,
	}
}

func (repo *EmployeeRepo) GetAllEmployee() []models.Employee {
	var employee []models.Employee
	err := repo.db.Find(&employee).Error

	if err != nil {
		return []models.Employee{}
	}
	return employee
}

func (repo *EmployeeRepo) GetEmployeeById(id uint) (models.Employee, error) {
	var employee models.Employee
	if err := repo.db.Where("ID = ?", id).First(&employee).Error; err != nil {
		return employee, err
	}
	return employee, nil
}

func (repo *EmployeeRepo) CreateEmployee(employee *models.Employee) error {
	err := repo.db.Create(employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *EmployeeRepo) UpdateEmployee(employee *models.Employee) error {
	err := repo.db.Save(employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *EmployeeRepo) DeleteEmployee(id uint) error {
	var employee models.Employee
	if err := repo.db.Where("ID = ?", id).Delete(&employee).Error; err != nil {
		return err
	}
	return nil
}
