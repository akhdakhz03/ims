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

type ProductInController struct{}

var productInObject = new(models.Model)
var productInParam parameterforms.Pagination
var ProductInTransactionParam parameterforms.ProductIn

func (w *ProductInController) GetBarangMasuk() gin.HandlerFunc {
	return func(c *gin.Context) {
		productInParam.Limit, _ = strconv.Atoi(c.Query("limit"))
		productInParam.Offset, _ = strconv.Atoi(c.Query("offset"))
		productInParam.Offset *= productInParam.Limit
		r := productInObject.GetBarangMasuk(productInParam.Limit, productInParam.Offset)
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

func (w *ProductInController) InsertBarangMasuk() gin.HandlerFunc {
	return func(c *gin.Context) {
		ProductInTransactionParam.Time = time.Now()
		ProductInTransactionParam.SKU = c.PostForm("sku")
		ProductInTransactionParam.ProductName = productInObject.GetNamaBarang(ProductInTransactionParam.SKU).ItemName
		ProductInTransactionParam.TotalStock, _ = strconv.Atoi(c.PostForm("jumlah_pemesanan"))
		ProductInTransactionParam.ActualStock, _ = strconv.Atoi(c.PostForm("jumlah_diterima"))
		ProductInTransactionParam.PricePerItem, _ = strconv.ParseFloat(c.PostForm("harga_beli"), 64)
		ProductInTransactionParam.TotalPrice = float64(ProductInTransactionParam.TotalStock) * ProductInTransactionParam.PricePerItem
		ProductInTransactionParam.Kwitansi = c.PostForm("nomor_kwitansi")
		ProductInTransactionParam.Remark = c.PostForm("catatan")

		var productIn tableform.ProductIn
		productIn.Time = ProductInTransactionParam.Time
		productIn.SKU = ProductInTransactionParam.SKU
		productIn.ProductName = ProductInTransactionParam.ProductName
		productIn.TotalStock = ProductInTransactionParam.TotalStock
		productIn.ActualStock = ProductInTransactionParam.ActualStock
		productIn.PricePerItem = ProductInTransactionParam.PricePerItem
		productIn.TotalPrice = ProductInTransactionParam.TotalPrice
		productIn.Kwitansi = ProductInTransactionParam.Kwitansi
		productIn.Remark = ProductInTransactionParam.Remark

		id := productInObject.InsertBarangMasuk(productIn)
		qty := productOutObject.GetQtyBarang(productOutTransactionParam.SKU)
		if id.Id == 0 {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Update Database Err",
			})

		} else {
			if qty.ItemName != "" {
				var currentStock int = qty.CurrentStock + ProductInTransactionParam.ActualStock
				productOutObject.UpdateStock(productOutTransactionParam.SKU, currentStock)
			} else {
				var currentStock int = ProductInTransactionParam.ActualStock
				productOutObject.UpdateStock(productOutTransactionParam.SKU, currentStock)
			}
			var history tableform.ProductHistory
			history.SKU = ProductInTransactionParam.SKU
			history.TransactionType = "in"
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

func (w *ProductInController) UpdateBarangMasuk() gin.HandlerFunc {
	return func(c *gin.Context) {
		ProductInTransactionParam.Time = time.Now()
		ProductInTransactionParam.SKU = c.PostForm("sku")
		ProductInTransactionParam.ProductName = productInObject.GetNamaBarang(ProductInTransactionParam.SKU).ItemName
		ProductInTransactionParam.TotalStock, _ = strconv.Atoi(c.PostForm("jumlah_pemesanan"))
		ProductInTransactionParam.ActualStock, _ = strconv.Atoi(c.PostForm("jumlah_diterima"))
		ProductInTransactionParam.PricePerItem, _ = strconv.ParseFloat(c.PostForm("harga_beli"), 64)
		ProductInTransactionParam.TotalPrice = float64(ProductInTransactionParam.TotalStock) * ProductInTransactionParam.PricePerItem
		ProductInTransactionParam.Kwitansi = c.PostForm("nomor_kwitansi")
		ProductInTransactionParam.Remark = c.PostForm("catatan")
		ProductInTransactionParam.Id, _ = strconv.Atoi(c.PostForm("id"))
		isExists := productInObject.GetBarangMasukById(ProductInTransactionParam.Id)
		if isExists.Id != 0 {
			var productIn tableform.ProductIn
			productIn.Id = ProductInTransactionParam.Id
			productIn.Time = ProductInTransactionParam.Time
			productIn.SKU = ProductInTransactionParam.SKU
			productIn.ProductName = ProductInTransactionParam.ProductName
			productIn.TotalStock = ProductInTransactionParam.TotalStock
			productIn.ActualStock = ProductInTransactionParam.ActualStock
			productIn.PricePerItem = ProductInTransactionParam.PricePerItem
			productIn.Kwitansi = ProductInTransactionParam.Kwitansi
			productIn.Remark = ProductInTransactionParam.Remark
			id := productInObject.UpdateBarangMasuk(productIn)
			qty := productOutObject.GetQtyBarang(productOutTransactionParam.SKU)
			if id.Id == 0 {
				c.JSON(500, gin.H{
					"status":  500,
					"message": "Update Database Err",
				})

			} else {
				if qty.ItemName != "" {
					var currentStock int = qty.CurrentStock + ProductInTransactionParam.ActualStock
					productOutObject.UpdateStock(productOutTransactionParam.SKU, currentStock)
				} else {
					var currentStock int = ProductInTransactionParam.ActualStock
					productOutObject.UpdateStock(productOutTransactionParam.SKU, currentStock)
				}
				var history tableform.ProductHistory
				history.SKU = ProductInTransactionParam.SKU
				history.TransactionType = "in"
				history.Date = time.Now()
				history.Stock = ProductInTransactionParam.ActualStock
				productOutObject.InsertHistory(history)
				c.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "success",
					"data":    id,
				})
			}
		} else {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Update Database Err",
			})
		}
	}
}

func (w *ProductInController) DeleteBarangMasuk() gin.HandlerFunc {
	return func(c *gin.Context) {
		ProductInTransactionParam.Id, _ = strconv.Atoi(c.PostForm("id"))
		r := productInObject.DeleteBarangMasuk(ProductInTransactionParam.Id)
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
