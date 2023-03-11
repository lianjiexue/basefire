//go:build ignore
// +build ignore

package main

import (
	"app/internal/models"
	"app/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	dsn := utils.GetDbDsn()
	gormdb, _ := gorm.Open(mysql.Open(dsn))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(models.User{}, models.Third{}, models.Role{}, models.RoleMenu{}, models.RoleAccess{}, models.Product{}, models.ProductSku{}, models.ProductImage{}, models.Order{}, models.OrderProducts{},
		models.OrderAddress{}, models.News{}, models.Menu{}, models.Manager{}, models.File{}, models.Config{}, models.ConfigGroup{}, models.Category{}, models.Cart{}, models.Banner{}, models.Article{}, models.Area{}, models.AreaProduct{}, models.Address{}, models.Access{})

	// Generate the code
	g.Execute()

}
