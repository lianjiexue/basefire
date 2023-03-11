package config

import (
	"github.com/spf13/viper"
	"log"
)

// NewConfig 初始化配置/*
func NewConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		panic("配置文件不存在")
	}
	//初始化数据库配置
	NewDatabase()
}

// GetDatabaseConfig 获取数据库的配置 /*
func GetDatabaseConfig() *Database {
	return database
}
