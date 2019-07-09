package controllers

import (
	"api-inventory/models"
	"api-inventory/utils/db"
	"api-inventory/utils/form/parameterforms"
	"api-inventory/utils/form/tableform"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductStockController struct{}

var productStockObject = new(models.Model)
var productStockParam parameterforms.Pagination
var productUpdate parameterforms.Products

func (w *ProductStockController) GetStockBarang() gin.HandlerFunc {
	return func(c *gin.Context) {
		productStockParam.Limit, _ = strconv.Atoi(c.Query("limit"))
		productStockParam.Offset, _ = strconv.Atoi(c.Query("offset"))
		productStockParam.Offset *= productStockParam.Limit
		r := productStockObject.GetStockBarang(productStockParam.Limit, productStockParam.Offset)
		if len(r) == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Data Not Found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
				"data":    r,
			})

		}
	}
}

func (w *ProductStockController) InsertStockBarang() gin.HandlerFunc {
	return func(c *gin.Context) {
		productUpdate.SKU = c.PostForm("sku")
		productUpdate.ItemName = c.PostForm("nama_produk")
		productUpdate.CurrentStock, _ = strconv.Atoi(c.PostForm("stok"))
		var product tableform.Products
		product.SKU = productUpdate.SKU
		product.CurrentStock = productUpdate.CurrentStock
		r := productStockObject.InsertProduct(product)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    r,
		})
	}
}

func (w *ProductStockController) UpdateStockBarang() gin.HandlerFunc {
	return func(c *gin.Context) {
		productUpdate.SKU = c.PostForm("sku")
		productUpdate.CurrentStock, _ = strconv.Atoi(c.PostForm("stok"))
		r := productStockObject.UpdateStock(productUpdate.SKU, productUpdate.CurrentStock)
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "success",
			"data":    r,
		})
	}
}

func (w *ProductStockController) CheckSqliteConnection() gin.HandlerFunc {
	return func(c *gin.Context) {
		db.InitSqlite()
	}
}
