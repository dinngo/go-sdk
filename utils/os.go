package utils

import "os"

func GetNullableEnv(key string) *string {
	env := os.Getenv(key)
	if len(env) > 0 {
		return &env
	}
	return nil
}

func GetEnvWithDefaultValue(key, defaultValue string) string {
	env := os.Getenv(key)
	if len(env) > 0 {
		return env
	}
	return defaultValue
}
