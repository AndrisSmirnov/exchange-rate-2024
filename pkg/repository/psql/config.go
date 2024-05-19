package pstgrs

import "time"

type Config struct {
	User                string
	Password            string
	Network             string
	Host                string
	Port                string
	DBName              string
	SSLMode             string
	MaxIdleCons         int
	RequestTimeout      time.Duration
	MaxIdleConnDuration time.Duration
	DriverName          string
}

func NewConfig(
	user string,
	password string,
	network string,
	host string,
	port string,
	dbName string,
	sslMode string,
	maxIdleCons int,
	requestTimeout time.Duration,
	maxIdleConnDuration time.Duration,
	driverName string,
) *Config {
	return &Config{
		User:                user,
		Password:            password,
		Network:             network,
		Host:                host,
		Port:                port,
		DBName:              dbName,
		SSLMode:             sslMode,
		MaxIdleCons:         maxIdleCons,
		RequestTimeout:      requestTimeout,
		MaxIdleConnDuration: maxIdleConnDuration,
		DriverName:          driverName,
	}
}
