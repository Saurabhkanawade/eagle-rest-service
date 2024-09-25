package dao

import (
	"context"
	"github.com/Saurabhkanawade/eagle-common-service/database"
	"github.com/gofrs/uuid"
	"github.com/saurabhkanawade/eagle-rest-service/internal/dbmodels"
	"github.com/saurabhkanawade/eagle-rest-service/pkg/models"
	"github.com/sirupsen/logrus"
)

var employees []*dbmodels.Employee

type EmployeeDao interface {
	Create(ctx context.Context, employee dbmodels.Employee) error
	GetAll(ctx context.Context) (dbmodels.EmployeeSlice, error)
	Get(ctx context.Context, employeeId uuid.UUID) (models.Employee, error)
	Update(ctx context.Context, employeeId uuid.UUID) (dbmodels.Employee, error)
	Delete(ctx context.Context, employeeId uuid.UUID) (dbmodels.Employee, error)
}

type EmployeeDaoImpl struct {
	conn database.DbConnection
}

func NewEmployeeDao(conn database.DbConnection) EmployeeDao {
	return &EmployeeDaoImpl{
		conn: conn,
	}
}

func (e EmployeeDaoImpl) Create(ctx context.Context, employee dbmodels.Employee) error {
	logrus.Info("employee_dao () - create employee")

	employees = append(employees, &employee)
	logrus.Infof("created the employee in db %v", employee)
	return nil
}

func (e EmployeeDaoImpl) GetAll(ctx context.Context) (dbmodels.EmployeeSlice, error) {
	logrus.Info("employee_dao () - get employees")
	return employees, nil
}

func (e EmployeeDaoImpl) Get(ctx context.Context, employeeId uuid.UUID) (models.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeDaoImpl) Update(ctx context.Context, employeeId uuid.UUID) (dbmodels.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeDaoImpl) Delete(ctx context.Context, employeeId uuid.UUID) (dbmodels.Employee, error) {
	//TODO implement me
	panic("implement me")
}
