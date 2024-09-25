package transport

import (
	"context"
	"encoding/json"
	commontransport "github.com/Saurabhkanawade/eagle-common-service/httptransport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/eagle-rest-service/internal/endpoint"
	"net/http"
)

func CreateEmployeeHandler(employeeEndpoint endpoint.EmployeeEndpoint, route *mux.Router) {
	route.Handle("/employee",
		httpTransport.NewServer(
			employeeEndpoint.CreateEmployee,
			decodeCreateEmployee,
			commontransport.EncodePostResponse,
		),
	).Methods(http.MethodPost)
}

func GetEmployeesHandler(employeeEndpoint endpoint.EmployeeEndpoint, route *mux.Router) {
	route.Handle("/employees",
		httpTransport.NewServer(
			employeeEndpoint.GetEmployees,
			decodeGetEmployees,
			commontransport.EncodeResponse,
		),
	).Methods(http.MethodGet)
}

func decodeCreateEmployee(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.AddEmployeeRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil

}

func decodeGetEmployees(ctx context.Context, r *http.Request) (interface{}, error) {
	return endpoint.GetEmployeeResponse{}, nil
}
