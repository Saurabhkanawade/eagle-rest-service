package main

import (
	"fmt"
	"github.com/Saurabhkanawade/eagle-common-service/config"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Welcome to the eagle-rest-service..........")

	config.LoadConfig(config.GetAppEnvLocation())

	if err := config.CheckRequiredVariables(); err != nil {
		logrus.Fatalf("%v", err)
	}

	StartServer()
}
