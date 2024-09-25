package models

import (
	"github.com/gofrs/uuid"
	"github.com/saurabhkanawade/eagle-rest-service/internal/dbmodels"
	"github.com/volatiletech/null/v8"
	"time"
)

type Employee struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	HireDate  time.Time `json:"hireDate"`
	Position  string    `json:"position"`
}

func ModelToDb(employee Employee) dbmodels.Employee {
	return dbmodels.Employee{
		UUID:      employee.ID.String(),
		Firstname: null.StringFrom(employee.FirstName),
		Lastname:  null.StringFrom(employee.LastName),
		Email:     null.StringFrom(employee.Email),
		Phone:     null.StringFrom(employee.Phone),
		HireDate:  employee.HireDate,
		Position:  null.StringFrom(employee.Position),
	}
}

func DbToModel(employee dbmodels.Employee) Employee {
	employeeId, _ := uuid.FromString(employee.UUID)
	return Employee{
		ID:        employeeId,
		FirstName: employee.Firstname.String,
		LastName:  employee.Lastname.String,
		Email:     employee.Email.String,
		Phone:     employee.Phone.String,
		HireDate:  employee.HireDate,
		Position:  employee.Position.String,
	}
}
