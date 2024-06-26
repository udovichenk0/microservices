package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnvironmentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}

func getEnvironmentValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return value
}

func GetPaymentServiceUrl() string {
	// return getEnvironmentValue("PAYMENT_SERVICE_URL")
	return "localhost:3001"
}
