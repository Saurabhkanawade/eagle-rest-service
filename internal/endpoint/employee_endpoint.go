package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/gofrs/uuid"
	"github.com/saurabhkanawade/eagle-rest-service/internal/service"
	"github.com/saurabhkanawade/eagle-rest-service/pkg/models"
	"github.com/sirupsen/logrus"
)

type EmployeeEndpoint struct {
	CreateEmployee endpoint.Endpoint
	GetEmployees   endpoint.Endpoint
	GetEmployee    endpoint.Endpoint
	UpdateEmployee endpoint.Endpoint
	DeleteEmployee endpoint.Endpoint
}

func MakeEmployeeEndpoints(s service.EmployeeService) EmployeeEndpoint {
	return EmployeeEndpoint{
		CreateEmployee: MakeCreateEmployeeEndpoint(s),
		GetEmployees:   MakeGetEmployeesEndpoint(s),
		GetEmployee:    MakeGetEmployeeEndpoint(s),
		UpdateEmployee: MakeUpdateEmployeeEndpoint(s),
		DeleteEmployee: MakeDeleteEmployeeEndpoint(s),
	}
}

type AddEmployeeRequest struct {
	Employee models.Employee `json:"employee"`
}

type AddEmployeeResponse struct{}

func MakeCreateEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		logrus.Info("employee_endpoint () - create employee endpoint")

		req, ok := request.(AddEmployeeRequest)
		if !ok {
			return nil, errors.New("error in endpoint")
		}

		logrus.Infof("Sending to service %v", req.Employee)
		err = s.Create(ctx, req.Employee)

		if err != nil {
			return nil, err
		}

		return AddEmployeeResponse{}, nil
	}
}

type GetEmployeeResponse struct {
	Employee []models.Employee `json:"employee,omitempty"`
}

func MakeGetEmployeesEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		logrus.Info("employee_endpoint () - get employees endpoint")

		employees, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		return employees, nil
	}
}

type GetEmployeeByIDParam struct {
	employeeID uuid.UUID
}

type GetEmployeeByIDResponse struct {
	employee models.Employee
}

func MakeGetEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		logrus.Info("employee_endpoint () - get employee endpoint")
		req, ok := request.(GetEmployeeByIDParam)
		if !ok {
			return nil, err
		}

		employee, err := s.Get(ctx, req.employeeID)

		return GetEmployeeByIDResponse{
			employee: employee,
		}, nil
	}
}

type UpdateEmployeeRequest struct {
	Employee models.Employee `json:"employee"`
}

type UpdateEmployeeResponse struct {
	Employee models.Employee `json:"employee"`
}

func MakeUpdateEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		logrus.Info("employee_endpoint () - update employee endpoint")
		req, ok := request.(UpdateEmployeeRequest)
		if !ok {
			return nil, err
		}
		updatedEmployee, err := s.Update(ctx, req.Employee, req.Employee.ID)

		return UpdateEmployeeResponse{
			Employee: updatedEmployee,
		}, nil
	}
}

type DeleteEmployeeRequestParam struct {
	employeeId uuid.UUID
}

type DeleteEmployeeResponse struct{}

func MakeDeleteEmployeeEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		logrus.Info("employee_endpoint () - delete employee endpoint")
		req, ok := request.(DeleteEmployeeRequestParam)
		if !ok {
			return nil, err
		}
		err = s.Delete(ctx, req.employeeId)
		if err != nil {
			return nil, err
		}
		return DeleteEmployeeResponse{}, nil
	}
}
