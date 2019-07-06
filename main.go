package main

import (
	"api-inventory/config"
	"api-inventory/controllers"
	"api-inventory/utils/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	ProductInController := new(controllers.ProductInController)
	ProductOutController := new(controllers.ProductOutController)
	ProductStockController := new(controllers.ProductStockController)
	confReader := new(config.ConfigReader)
	confReader.Read()
	db.InitMysql()
	r := gin.Default()
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	r.Use(cors.New(c))

	r.GET("/get_barang_masuk", ProductInController.GetBarangMasuk())
	r.GET("/get_barang_keluar", ProductOutController.GetBarangKeluar())
	r.GET("/get_stok_barang", ProductStockController.GetStockBarang())

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
	if viper.GetBool("secure") == false {
		r.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
	} else {
		r.RunTLS(viper.GetString("server.host")+":"+viper.GetString("server.port"), viper.GetString("crt_file"), viper.GetString("key_file"))
	}
}
