package main

import (
	"context"
	"github.com/Saurabhkanawade/eagle-common-service/config"
	"github.com/Saurabhkanawade/eagle-common-service/database"
	"github.com/Saurabhkanawade/eagle-common-service/server"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/eagle-rest-service/internal/dao"
	"github.com/saurabhkanawade/eagle-rest-service/internal/endpoint"
	"github.com/saurabhkanawade/eagle-rest-service/internal/service"
	"github.com/saurabhkanawade/eagle-rest-service/internal/transport"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

func StartServer() {
	DbConfig := database.DbConfig{
		Host:   viper.GetString("SERVER_PORT"),
		Port:   viper.GetInt("POSTGRES_PORT"),
		User:   viper.GetString("POSTGRES_USER"),
		Pass:   viper.GetString("POSTGRES_PASS"),
		DbName: viper.GetString("POSTGRES_DB"),
	}

	dbConn, err := database.InitDatabase(DbConfig)
	if err != nil {
		return
	}
	logrus.Infof("Connection successfull to db %v", &dbConn)

	ctx := context.Background()
	router := mux.NewRouter().StrictSlash(true)

	//dao
	employeeDao := dao.NewEmployeeDao(dbConn)

	//service
	employeeService := service.NewEmployeeService(employeeDao)

	//endpoint
	employeeEndpoint := endpoint.MakeEmployeeEndpoints(employeeService)

	//transport
	transport.CreateEmployeeHandler(employeeEndpoint, router)
	transport.GetEmployeesHandler(employeeEndpoint, router)

	err = server.StartServer(ctx, router,
		server.SetPort(config.GetServerPort()),
		server.SetReadTimeout(time.Duration(config.GetReadTimeout())*time.Second),
		server.SetWriteTimeout(time.Duration(config.GetWriteTimeout())*time.Second),
	)
	if err != nil {
		logrus.Errorf("error from http server: %v", err)
	}
}
