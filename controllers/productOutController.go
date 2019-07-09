package controllers

import (
	"api-inventory/models"
	"api-inventory/utils/form/parameterforms"
	"api-inventory/utils/form/tableform"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductOutController struct{}

var productOutObject = new(models.Model)
var productOutParam parameterforms.Pagination
var productOutTransactionParam parameterforms.ProductOut

func (w *ProductOutController) GetBarangKeluar() gin.HandlerFunc {
	return func(c *gin.Context) {
		productOutParam.Limit, _ = strconv.Atoi(c.Query("limit"))
		productOutParam.Offset, _ = strconv.Atoi(c.Query("offset"))
		productOutParam.Offset *= productOutParam.Limit
		r := productOutObject.GetBarangKeluar(productOutParam.Limit, productOutParam.Offset)
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

func (w *ProductOutController) InsertBarangKeluar() gin.HandlerFunc {
	return func(c *gin.Context) {
		productOutTransactionParam.Time = time.Now()
		productOutTransactionParam.SKU = c.PostForm("sku")
		productOutTransactionParam.ProductName = productOutObject.GetNamaBarang(productOutTransactionParam.SKU).ItemName
		productOutTransactionParam.Qty, _ = strconv.Atoi(c.PostForm("jumlah_keluar"))
		productOutTransactionParam.SellPrice, _ = strconv.ParseFloat(c.PostForm("harga_jual"), 64)
		productOutTransactionParam.TotalPrice = productOutTransactionParam.SellPrice * float64(productOutTransactionParam.Qty)
		productOutTransactionParam.Remark = c.PostForm("catatan")

		var productOut tableform.ProductOut
		productOut.Time = productOutTransactionParam.Time
		productOut.ProductName = productOutTransactionParam.ProductName
		productOut.Qty = productOutTransactionParam.Qty
		productOut.SellPrice = productOutTransactionParam.SellPrice
		productOut.TotalPrice = productOutTransactionParam.TotalPrice
		productOut.Remark = productOutTransactionParam.Remark
		productOut.SKU = productOutTransactionParam.SKU

		//if stock < qty maka error
		qty := productOutObject.GetQtyBarang(productOutTransactionParam.SKU)
		if qty.ItemName == "" {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Product Not Found",
			})
		} else if qty.CurrentStock < productOutTransactionParam.Qty {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Stock Kurang Dari Jumlah",
			})
		} else {

			id := productOutObject.InsertBarangKeluar(productOut)

			if id.Id == 0 {
				c.JSON(500, gin.H{
					"status":  500,
					"message": "Update Database Err",
				})
				c.Abort()
				return
			} else {
				var currentStock int = qty.CurrentStock - productOutTransactionParam.Qty
				productOutObject.UpdateStock(productOutTransactionParam.SKU, currentStock)
				var history tableform.ProductHistory
				history.SKU = ProductInTransactionParam.SKU
				history.TransactionType = "out"
				history.Date = time.Now()
				history.Stock = ProductInTransactionParam.ActualStock
				productOutObject.InsertHistory(history)
				c.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "success",
					"data":    id,
				})

			}

		}

	}
}

func (w *ProductOutController) UpdateBarangKeluar() gin.HandlerFunc {
	return func(c *gin.Context) {
		productOutTransactionParam.Id, _ = strconv.Atoi(c.PostForm("id"))
		productOutTransactionParam.Time = time.Now()
		productOutTransactionParam.SKU = c.PostForm("sku")
		productOutTransactionParam.ProductName = productOutObject.GetNamaBarang(productOutTransactionParam.SKU).ItemName
		productOutTransactionParam.Qty, _ = strconv.Atoi(c.PostForm("jumlah_keluar"))
		productOutTransactionParam.SellPrice, _ = strconv.ParseFloat(c.PostForm("harga_jual"), 64)
		productOutTransactionParam.TotalPrice = productOutTransactionParam.SellPrice * float64(productOutTransactionParam.Qty)
		productOutTransactionParam.Remark = c.PostForm("catatan")

		var productOut tableform.ProductOut
		productOut.Id = productOutTransactionParam.Id
		productOut.Time = productOutTransactionParam.Time
		productOut.ProductName = productOutTransactionParam.ProductName
		productOut.Qty = productOutTransactionParam.Qty
		productOut.SellPrice = productOutTransactionParam.SellPrice
		productOut.TotalPrice = productOutTransactionParam.TotalPrice
		productOut.Remark = productOutTransactionParam.Remark
		productOut.SKU = productOutTransactionParam.SKU

		isExists := productOutObject.GetBarangKeluarById(productOutTransactionParam.Id)
		//if stock < qty maka error
		if isExists.Id != 0 {

			qty := productOutObject.GetQtyBarang(productOutTransactionParam.SKU)
			if qty.ItemName == "" {
				c.JSON(500, gin.H{
					"status":  500,
					"message": "Product Not Found",
				})
			} else if qty.CurrentStock < productOutTransactionParam.Qty {
				c.JSON(500, gin.H{
					"status":  500,
					"message": "Stock Kurang Dari Jumlah",
				})
			} else {

				id := productOutObject.UpdateBarangKeluar(productOut)

				if id.Id == 0 {
					c.JSON(500, gin.H{
						"status":  500,
						"message": "Update Database Err",
					})
					c.Abort()
					return
				} else {
					var currentStock int = qty.CurrentStock - productOutTransactionParam.Qty
					productOutObject.UpdateStock(productOutTransactionParam.SKU, currentStock)
					var history tableform.ProductHistory
					history.SKU = ProductInTransactionParam.SKU
					history.TransactionType = "out"
					history.Date = time.Now()
					history.Stock = ProductInTransactionParam.ActualStock
					productOutObject.InsertHistory(history)
					c.JSON(http.StatusOK, gin.H{
						"status":  http.StatusOK,
						"message": "success",
						"data":    id,
					})
				}
			}
		} else {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Update Database Err",
			})
		}
	}
}

func (w *ProductOutController) DeleteBarangKeluar() gin.HandlerFunc {
	return func(c *gin.Context) {
		productOutTransactionParam.Id, _ = strconv.Atoi(c.PostForm("id"))
		r := productOutObject.DeleteBarangKeluar(productOutTransactionParam.Id)
		if r == true {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
			})
		} else {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Error Deleting Data",
			})
		}
	}
}
