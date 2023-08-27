package database

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
)

type CustomIP net.IP

func (ip *CustomIP) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal IP value:", value))
	}

	ipnet := net.ParseIP(str)

	if ipnet != nil {
		*ip = CustomIP(ipnet)
	}

	return errors.New(fmt.Sprint("Non valid text represensation of an IP address:", value))
}

func (ip CustomIP) Value() (driver.Value, error) {
	return net.IP(ip).String(), nil
}
