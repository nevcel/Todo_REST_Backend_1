// Package configuration contains the logic for accessing the configuration
package configuration

import (
	"github.com/joho/godotenv"
	"strconv"
)

const EnvFile = ".env"
const RepositoryModeKeyName = "REPOSITORY_MODE"
const PortKeyName = "PORT"
const MemoryRepository = "mem"
const CsvFileRepository = "csv"
const RepositoryModeDefault = MemoryRepository
const PortDefault = 8080

// GetRepositoryMode returns the configured repositories mode
func GetRepositoryMode() (string, error) {
	configMap, err := GetConfiguration()
	if err != nil {
		return "", err
	}

	repositoryMode := configMap[RepositoryModeKeyName]
	if repositoryMode == "" {
		repositoryMode = RepositoryModeDefault
	}

	return repositoryMode, nil
}

// GetConfiguration returns a map containing the configurations
func GetConfiguration() (map[string]string, error) {
	return godotenv.Read(EnvFile)
}

// GetBackendHostUrl returns the configured backend url
func GetBackendHostUrl() (string, error) {
	configMap, err := GetConfiguration()
	if err != nil {
		return "", err
	}

	port := configMap[PortKeyName]
	if port == "" {
		port = strconv.Itoa(PortDefault)
	}

	backendHostUrl := ":" + port

	return backendHostUrl, nil
}
