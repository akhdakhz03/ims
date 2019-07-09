package tableform

import "time"

//tableform is use as connection between model and database

type ProductIn struct {
	Id           int       `gorm:"column:Id;PRIMARY_KEY"`
	Time         time.Time `gorm:"column:Time"`
	SKU          string    `gorm:"column:SKU"`
	ProductName  string    `gorm:"column:ProductName"`
	TotalStock   int       `gorm:"column:TotalStock"`
	ActualStock  int       `gorm:"column:ActualStock"`
	PricePerItem float64   `gorm:"column:PricePerItem"`
	TotalPrice   float64   `gorm:"column:TotalPrice"`
	Kwitansi     string    `gorm:"column:Kwitansi"`
	Remark       string    `gorm:"column:Remark"`
}

type ProductOut struct {
	Id          int       `gorm:"column:Id;PRIMARY_KEY"`
	Time        time.Time `gorm:"column:Time"`
	ProductName string    `gorm:"column:ProductName"`
	Qty         int       `gorm:"column:Qty"`
	SellPrice   float64   `gorm:"column:SellPrice"`
	TotalPrice  float64   `gorm:"column:TotalPrice"`
	Remark      string    `gorm:"column:Remark"`
	SKU         string    `gorm:"column:SKU"`
}

type Products struct {
	SKU          string `gorm:"column:SKU;PRIMARY_KEY"`
	ItemName     string `gorm:"column:ItemName"`
	CurrentStock int    `gorm:"column:CurrentStock"`
}

type ProductHistory struct {
	SKU             string    `gorm:"column:SKU"`
	TransactionType string    `gorm:"column:TransactionType"`
	Date            time.Time `gorm:"column:Date"`
	Stock           int       `gorm:"column:Stock"`
}

type LaporanNilaiBarang struct {
	SKU          string  `gorm:"column:SKU;PRIMARY_KEY"`
	NamaItem     string  `gorm:"column:ItemName"`
	Jumlah       int     `gorm:"column:CurrentStock"`
	HargaBeliAvg float64 `gorm:"column:avgprice"`
	Total        float64 `gorm:"column:total"`
}

type LaporanPenjualan struct {
	IdPesanan  string `gorm:"column:IdPesanan"`
	SKU        string `gorm:"column:SKU"`
	Waktu      string `gorm:"column:Waktu"`
	NamaBarang string `gorm:"column:NamaBarang"`
	Jumlah     string `gorm:"column:Jumlah"`
	HargaJual  string `gorm:"column:HargaJual"`
	Total      string `gorm:"column:Total"`
	HargaBeli  string `gorm:"column:HargaBeli"`
	Laba       string `gorm:"column:Laba"`
}
