package Core

import (
	"github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/mysql"
)
var DB *gorm.DB

func init(){
	db, err := gorm.Open("mysql", "root:@/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = db
	//defer db.Close()
}
