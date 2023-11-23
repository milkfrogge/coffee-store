package config

import (
	"errors"
	"fmt"
	"net"
	"os"
)

const (
	defaultHostEnvName        = "PRODUCT_SERVER_HOST"
	defaultPortEnvName        = "PRODUCT_SERVER_PORT"
	defaultDbHostEnvName      = "PRODUCT_DB_HOST"
	defaultDbPortEnvName      = "PRODUCT_DB_PORT"
	defaultDbUsernameEnvName  = "PRODUCT_DB_USERNAME"
	defaultDbPasswordEnvName  = "PRODUCT_DB_PASSWORD"
	defaultDbNameEnvName      = "PRODUCT_DB_NAME"
	defaultJaegerHostEnvName  = "JAEGER_HOST"
	defaultJaegerPortEnvName  = "JAEGER_PORT"
	defaultGraylogHostEnvName = "GRAYLOG_HOST"
	defaultGraylogPortEnvName = "GRAYLOG_PORT"
)

type ProductConfig struct {
	host        string
	port        string
	dbHost      string
	dbPort      string
	JaegerHost  string
	JaegerPort  string
	GraylogHost string
	GraylogPort string
	dbUsername  string
	dbPassword  string
	dbName      string
}

func NewProductConfig() (*ProductConfig, error) {
	host := os.Getenv(defaultHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(defaultPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	dbHost := os.Getenv(defaultDbHostEnvName)
	if len(dbHost) == 0 {
		return nil, errors.New("dbHost not found")
	}

	dbPort := os.Getenv(defaultDbPortEnvName)
	if len(dbPort) == 0 {
		return nil, errors.New("dbPort not found")
	}

	dbName := os.Getenv(defaultDbNameEnvName)
	if len(dbName) == 0 {
		return nil, errors.New("dbName not found")
	}

	dbUser := os.Getenv(defaultDbUsernameEnvName)
	if len(dbUser) == 0 {
		return nil, errors.New("dbUser not found")
	}

	dbPass := os.Getenv(defaultDbPasswordEnvName)
	if len(dbPass) == 0 {
		return nil, errors.New("dbPass not found")
	}

	jaegerHost := os.Getenv(defaultJaegerHostEnvName)
	if len(jaegerHost) == 0 {
		return nil, errors.New("jaegerHost not found")
	}

	jaegerPort := os.Getenv(defaultJaegerPortEnvName)
	if len(jaegerPort) == 0 {
		return nil, errors.New("jaegerPort not found")
	}

	graylogHost := os.Getenv(defaultGraylogHostEnvName)
	if len(graylogHost) == 0 {
		return nil, errors.New("graylogHost not found")
	}

	graylogPort := os.Getenv(defaultGraylogPortEnvName)
	if len(graylogPort) == 0 {
		return nil, errors.New("graylogPort not found")
	}

	return &ProductConfig{
		host:        host,
		port:        port,
		dbHost:      dbHost,
		dbPort:      dbPort,
		dbUsername:  dbName,
		dbPassword:  dbPass,
		dbName:      dbName,
		JaegerHost:  jaegerHost,
		JaegerPort:  jaegerPort,
		GraylogHost: graylogHost,
		GraylogPort: graylogPort,
	}, nil
}

func (p *ProductConfig) GetPostgresDsn() string {
	//"postgres://product:product@localhost:5432/product"
	return fmt.Sprintf("postgresql://%s:%s@%s/%s", p.dbUsername, p.dbPassword, net.JoinHostPort(p.dbHost, p.dbPort), p.dbName)
}

func (p *ProductConfig) Address() string {
	//"postgres://product:product@localhost:5432/product"
	return net.JoinHostPort(p.host, p.port)
}
