package repo

import (
	"app/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewDb() {
	var err error
	dns := utils.GetDbDsn()
	fmt.Println(dns)
	db, err = gorm.Open(mysql.Open(dns))
	if err != nil {
		//fmt.Println("数据库链接失败")
		panic("数据库链接失败")
	}
}
