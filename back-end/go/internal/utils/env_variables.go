package utils

import (
    "fmt"
    "os"
)

func GetEnvVariables(envVariablesMap map[string]string) error {
    var key string

    for key, _ = range envVariablesMap {
        envVariablesMap[key] = os.Getenv(key)

        if envVariablesMap[key] == "" {
            return fmt.Errorf("Failed to read the %s environment variable: it isn't set", key)
        }
    }

    return nil
}
