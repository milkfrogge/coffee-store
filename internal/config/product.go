package config

import (
	"errors"
	"fmt"
	"net"
	"os"
)

const (
	defaultHostEnvName       = "PRODUCT_SERVER_HOST"
	defaultPortEnvName       = "PRODUCT_SERVER_PORT"
	defaultDbHostEnvName     = "PRODUCT_DB_HOST"
	defaultDbPortEnvName     = "PRODUCT_DB_PORT"
	defaultDbUsernameEnvName = "PRODUCT_DB_USERNAME"
	defaultDbPasswordEnvName = "PRODUCT_DB_PASSWORD"
	defaultDbNameEnvName     = "PRODUCT_DB_NAME"
)

type ProductConfig struct {
	host       string
	port       string
	dbHost     string
	dbPort     string
	dbUsername string
	dbPassword string
	dbName     string
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

	return &ProductConfig{
		host:       host,
		port:       port,
		dbHost:     dbHost,
		dbPort:     dbPort,
		dbUsername: dbName,
		dbPassword: dbPass,
		dbName:     dbName,
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
