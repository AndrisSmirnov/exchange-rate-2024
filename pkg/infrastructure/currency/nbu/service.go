package nbu

import "exchange_rate/pkg/packages/errors"

type NBU struct {
	cnf *Config
}

func NewServiceNBU() (*NBU, *errors.Error) {
	cnf, err := NewConfig()
	if err != nil {
		return nil, err
	}

	if err := cnf.validate(); err != nil {
		return nil, err
	}

	return &NBU{
		cnf: cnf,
	}, nil
}
