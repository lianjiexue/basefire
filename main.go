package main

import (
	"app/bootstrap"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func main() {

	app := bootstrap.Application{Name: "App", Port: ":8080"}
	app.Start()
}
