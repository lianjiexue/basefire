package utils

import (
	"app/config"
	"fmt"
)

// GetDbDsn 获取数据库的dsn方便其它地方使用
func GetDbDsn() string {
	dbConfig := config.GetDatabaseConfig()
	fmt.Println(dbConfig)
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DbUsername,
		dbConfig.DbPassword,
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbDatabase)
}
