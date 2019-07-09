package controllers

import (
	"api-inventory/models"
	"api-inventory/utils/form/parameterforms"
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportController struct{}

var reportObject = new(models.Model)
var dateRange parameterforms.DateRange

func (w *ReportController) GetAveragePrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := reportObject.GetLaporanNilaiBarangDetail()
		var jumlah_total_barang int
		var total_nilai float64
		for _, dataR := range r {
			jumlah_total_barang += dataR.Jumlah
			total_nilai += dataR.Total
			//log.Println(dataR.CurrentStock)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":              http.StatusOK,
			"message":             "success",
			"tanggal_cetak":       "",
			"jumlah_sku":          len(r),
			"jumlah_total_barang": jumlah_total_barang,
			"total_nilai":         total_nilai,
			"data":                r,
		})
	}
}

func (w *ReportController) GetSellingReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		dateRange.DateStart = c.Query("date_start")
		dateRange.DateEnd = c.Query("date_end")
		r := reportObject.GetLaporanPenjualan(dateRange.DateStart, dateRange.DateEnd)
		var total_omzet float64 = 0
		var total_barang float64 = 0
		var total_laba_kotor float64 = 0
		var total_penjualan int64 = reportObject.GetTotalPenjualan(dateRange.DateStart, dateRange.DateEnd)
		for _, dataR := range r {
			omzet, _ := strconv.ParseFloat(dataR.Total, 64)
			jml, _ := strconv.ParseFloat(dataR.Jumlah, 64)
			labaKotor, _ := strconv.ParseFloat(dataR.Laba, 64)
			total_omzet += omzet
			total_barang += jml
			total_laba_kotor += labaKotor
		}

		c.JSON(http.StatusOK, gin.H{
			"status":           http.StatusOK,
			"message":          "success",
			"tanggal_cetak":    dateRange.DateStart + " - " + dateRange.DateEnd,
			"total_omzet":      total_omzet,
			"total_laba_kotor": total_laba_kotor,
			"total_penjualan":  total_penjualan,
			"total_barang":     total_barang,
			"data":             r,
		})
	}
}

func (w *ReportController) DownloadCSV() gin.HandlerFunc {
	return func(c *gin.Context) {
		var fileName string
		var fullUrl string
		//var baseUrl string
		//baseUrl = "."
		fileName = c.PostForm("filename")
		fullUrl = "./csv/laporan_" + c.PostForm("filename") + ".csv"

		file, err := os.OpenFile(fullUrl, os.O_CREATE|os.O_WRONLY, 0777)
		defer file.Close()
		if err != nil {
			os.Exit(1)
		}
		if fileName == "penjualan" {
			r := reportObject.GetAllLaporanPenjualan()
			csvWriter := csv.NewWriter(file)

			header := []string{"IdPesanan", "SKU", "NamaBarang", "Jumlah", "HargaBeli", "HargaJual", "Laba"}
			csvWriter.Write(header)
			for _, dataLaporan := range r {
				log.Println(dataLaporan)
				x := []string{dataLaporan.IdPesanan, dataLaporan.SKU, dataLaporan.NamaBarang, dataLaporan.Jumlah, dataLaporan.HargaBeli, dataLaporan.HargaJual, dataLaporan.Laba}
				csvWriter.Write(x)
			}
			csvWriter.Flush()
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
				"file":    "localhost:8080/csv/laporan_" + fileName + ".csv",
			})
		} else if fileName == "nilai_barang" {
			r := reportObject.GetLaporanNilaiBarangDetail()
			csvWriter := csv.NewWriter(file)

			header := []string{"SKU", "NamaBarang", "HargaBeli", "Jumlah", "Total"}
			csvWriter.Write(header)
			for _, dataLaporan := range r {
				log.Println(dataLaporan)
				x := []string{dataLaporan.SKU, dataLaporan.NamaItem, strconv.FormatFloat(dataLaporan.HargaBeliAvg, 'f', 0, 64), strconv.Itoa(dataLaporan.Jumlah), strconv.FormatFloat(dataLaporan.Total, 'f', 0, 64)}
				csvWriter.Write(x)
			}
			csvWriter.Flush()
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "success",
				"file":    "localhost:8080/csv/laporan_" + fileName + ".csv",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Data Not Found",
			})
		}
	}
}
