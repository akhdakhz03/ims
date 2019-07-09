package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

//db.go is use to open a database connection, in this case we will use mysql as database
var db *gorm.DB

func InitMysql() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", viper.GetString("database.mysql.user")+":"+viper.GetString("database.mysql.password")+"@tcp("+viper.GetString("database.mysql.host")+":"+viper.GetString("database.mysql.port")+")/"+viper.GetString("database.mysql.dbname")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")

	}
}

// GetDB is
func GetDB() *gorm.DB {
	return db
}

func InitSqlite() {
	var err error
	db, err = gorm.Open("sqlite3", "./inventory.sqlite")
	if err != nil {
		log.Println(err)
		panic("failed to connect database, sqlite")
	} else {
		log.Println("success connect sqlite")
	}
}
