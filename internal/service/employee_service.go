package service

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/saurabhkanawade/eagle-rest-service/internal/dao"
	"github.com/saurabhkanawade/eagle-rest-service/internal/dbmodels"
	"github.com/saurabhkanawade/eagle-rest-service/pkg/models"
	"github.com/sirupsen/logrus"
)

type EmployeeService interface {
	Create(ctx context.Context, employee models.Employee) error
	GetAll(ctx context.Context) ([]models.Employee, error)
	Get(ctx context.Context, employeeId uuid.UUID) (models.Employee, error)
	Update(ctx context.Context, updatedEmployee models.Employee, employeeId uuid.UUID) (models.Employee, error)
	Delete(ctx context.Context, employeeId uuid.UUID) error
}

type EmployeeServiceImpl struct {
	dao dao.EmployeeDao
}

func NewEmployeeService(employeeDao dao.EmployeeDao) EmployeeService {
	return &EmployeeServiceImpl{
		dao: employeeDao,
	}
}

func (e EmployeeServiceImpl) Create(ctx context.Context, emp models.Employee) error {
	logrus.Info("employee_service () - create employee")
	var employee dbmodels.Employee

	employee = models.ModelToDb(emp)
	err := e.dao.Create(ctx, employee)
	if err != nil {
		return err
	}

	return nil
}

func (e EmployeeServiceImpl) GetAll(ctx context.Context) ([]models.Employee, error) {
	logrus.Info("employee_service () - get employees")

	var employeeModelResponse []models.Employee

	employees, err := e.dao.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, emp := range employees {
		employeeModel := models.DbToModel(*emp)
		employeeModelResponse = append(employeeModelResponse, employeeModel)
	}

	return employeeModelResponse, nil
}

func (e EmployeeServiceImpl) Get(ctx context.Context, employeeId uuid.UUID) (models.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeServiceImpl) Update(ctx context.Context, updateEmployee models.Employee, employeeId uuid.UUID) (models.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeServiceImpl) Delete(ctx context.Context, employeeId uuid.UUID) error {
	panic("implement me")
}
