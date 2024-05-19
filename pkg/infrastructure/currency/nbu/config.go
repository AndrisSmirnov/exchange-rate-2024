package nbu

import (
	"net/url"

	"exchange_rate/pkg/packages/errors"
	"exchange_rate/pkg/utils"
)

type Config struct {
	BaseLink *url.URL
}

func (c *Config) validate() *errors.Error {
	if c.BaseLink == nil {
		return errConfigNotValid
	}

	return nil
}

func NewConfig() (*Config, *errors.Error) {
	nbuBaseLink, err := utils.TryGetEnv[string]("NBU_EXCHANGE_BASE_LINK")
	if err != nil {
		return nil, newErrorNoEnvVar("NBU_EXCHANGE_BASE_LINK")
	}

	u, errC := url.Parse(nbuBaseLink)
	if errC != nil {
		return nil, errToCreateConfig(err)
	}

	return &Config{
		BaseLink: u,
	}, nil
}
