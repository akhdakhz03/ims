package db

import "github.com/spf13/viper"

//db.go is use to open a database connection, in this case we will use mysql as database

func InitMysql() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", viper.GetString("database.mysql.user")+":"+viper.GetString("database.mysql.password")+"@tcp("+viper.GetString("database.mysql.host")+":"+viper.GetString("database.mysql.port")+")/"+viper.GetString("database.mysql.dbname")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
}
