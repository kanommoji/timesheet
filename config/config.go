package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Username string `yamal:"username"`
	Password string `yamal:"password"`
	Host     string `yamal:"host"`
	Port     string `yamal:"port"`
	Database string `yamal:"database"`
}

func (databaseConfig DatabaseConfig) GetURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Database)
}

func SetupDatabaseConfig(databaseConfigPath string) (DatabaseConfig, error) {
	var databaseConfig DatabaseConfig

	databaseConfigFile, err := ioutil.ReadFile(databaseConfigPath)
	if err != nil {
		return databaseConfig, err
	}

	err = yaml.Unmarshal(databaseConfigFile, &databaseConfig)
	if err != nil {
		return databaseConfig, err
	}

	return databaseConfig, err
}
