package utils

import (
	"os"
	"strconv"
	"strings"

	"exchange_rate/pkg/packages/errors"
)

type env interface {
	~string | ~int | ~float64 | ~float32 | ~[]string | ~bool
}

func TryGetEnvDefault[T env](env string, val T) T {
	res, err := getEnv[T](env)
	if err != nil {
		return val
	}

	return res
}

func TryGetEnv[T env](env string) (T, *errors.Error) {
	res, err := getEnv[T](env)
	if err != nil {
		return res, err
	}

	return res, nil
}

func getEnv[T env](env string) (T, *errors.Error) {
	var stdType T

	envVal := os.Getenv(env)
	if envVal == "" {
		return stdType, newErrNoEnv(env)
	}

	var defValue T

	switch (any(stdType)).(type) {
	case string:
		val, ok := any(envVal).(T)
		if !ok {
			return defValue, errorUnexpectedType
		}
		return val, nil
	case []string:
		val, ok := any(strings.Split(envVal, ",")).(T)
		if !ok {
			return defValue, errorUnexpectedType
		}
		return val, nil
	case int:
		res, err := strconv.Atoi(envVal)
		if err != nil {
			return stdType, newErrParse(env, "int", err.Error())
		}

		val, ok := any(res).(T)
		if !ok {
			return defValue, errorUnexpectedType
		}
		return val, nil

	case float64:
		res, err := strconv.ParseFloat(envVal, 64)
		if err != nil {
			return stdType, newErrParse(env, "float64", err.Error())
		}

		val, ok := any(res).(T)
		if !ok {
			return defValue, errorUnexpectedType
		}
		return val, nil

	case float32:
		res, err := strconv.ParseFloat(envVal, 32)
		if err != nil {
			return stdType, newErrParse(env, "float32", err.Error())
		}

		val, ok := any(res).(T)
		if !ok {
			return defValue, errorUnexpectedType
		}
		return val, nil

	case bool:
		res, err := strconv.ParseBool(envVal)
		if err != nil {
			return stdType, newErrParse(env, "bool", err.Error())
		}

		val, ok := any(res).(T)
		if !ok {
			return defValue, errorUnexpectedType
		}
		return val, nil

	default:
		return stdType, errorParseEnv
	}
}
