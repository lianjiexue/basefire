package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var database *Database

type Database struct {
	DbPrefix     string
	DbConnection string
	DbHost       string
	DbDatabase   string
	DbPort       string
	DbUsername   string
	DbPassword   string
}

func init() {
	database = NewDatabase()
}

func NewDatabase() *Database {
	database = new(Database)
	database.DbConnection = fmt.Sprintf("%s", viper.Get("DB_CONNECTION"))
	database.DbHost = fmt.Sprintf("%s", viper.Get("DB_HOST"))
	database.DbDatabase = fmt.Sprintf("%s", viper.Get("DB_DATABASE"))
	database.DbPort = fmt.Sprintf("%s", viper.Get("DB_PORT"))
	database.DbUsername = fmt.Sprintf("%s", viper.Get("DB_USERNAME"))
	database.DbPassword = fmt.Sprintf("%s", viper.Get("DB_PASSWORD"))
	return database
}
