package env

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func Required(name string, fns ...func(string) error) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Sprintf("missing required environment variable: %s", name))
	}

	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Sprintf("invalid value for environment variable: %s: %s", name, err))
		}
	}

	return v
}

func RequiredInt(name string, fns ...func(int) error) int {
	value := Required(name)

	v, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("invalid integer value for environment variable: %s", name))
	}

	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Sprintf("invalid integer value for environment variable: %s: %s", name, err))
		}
	}

	return v
}

func RequiredInt32(name string, fns ...func(int32) error) int32 {
	value := Required(name)

	v64, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("invalid integer value for environment variable: %s", name))
	}

	v := int32(v64)
	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Sprintf("invalid integer value for environment variable: %s: %s", name, err))
		}
	}

	return v
}

func RequiredInt64(name string, fns ...func(int64) error) int64 {
	value := Required(name)

	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid integer value for environment variable: %s", name))
	}

	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Sprintf("invalid integer value for environment variable: %s: %s", name, err))
		}
	}

	return v
}

func RequiredFloat32(name string, fns ...func(float32) error) float32 {
	value := Required(name)

	v64, err := strconv.ParseFloat(value, 32)
	if err != nil {
		panic(fmt.Sprintf("invalid float value for environment variable: %s", name))
	}

	v := float32(v64)
	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Sprintf("invalid float value for environment variable: %s: %s", name, err))
		}
	}

	return v
}

func RequiredFloat64(name string, fns ...func(float64) error) float64 {
	value := Required(name)

	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid float value for environment variable: %s", name))
	}

	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Sprintf("invalid float value for environment variable: %s: %s", name, err))
		}
	}

	return v
}

func RequiredBool(name string) bool {
	value := Required(name)

	v, err := strconv.ParseBool(value)
	if err != nil {
		panic(fmt.Sprintf("invalid boolean value for environment variable: %s", name))
	}

	return v
}

func OptionalBool(name string) bool {
	value, ok := os.LookupEnv(name)
	if !ok {
		return false
	}

	v, err := strconv.ParseBool(value)
	if err != nil {
		panic(fmt.Sprintf("invalid boolean value for environment variable: %s", name))
	}

	return v
}

func RequiredURL(name string, fns ...func(string) error) *url.URL {
	v, err := url.Parse(Required(name, fns...))
	if err != nil {
		panic(fmt.Sprintf("invalid URL value for environment variable: %s", name))
	}

	return v
}

func OptionalURL(name string, fns ...func(string) error) *url.URL {
	v, ok := os.LookupEnv(name)
	if !ok {
		return nil
	}

	u, err := url.Parse(v)
	if err != nil {
		panic(fmt.Sprintf("invalid URL value for environment variable: %s", name))
	}

	return u
}
