package model
 
import (
	"fmt"
 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)
 
var DB *gorm.DB
 
func ConnectDatabase() {
	dsn := "aprh8956_dbceria:dbceriangentot@tcp(103.247.8.36:3306)/aprh8956_dbceria?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	DB = db
}