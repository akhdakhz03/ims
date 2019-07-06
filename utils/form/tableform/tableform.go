package tableform

import "time"

//tableform is use as connection between model and database

type ProductIn struct {
	Id           int       `gorm:"column:Id"`
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
	Id          int       `gorm:"column:Id"`
	Time        time.Time `gorm:"column:Time"`
	ProductName string    `gorm:"column:ProductName"`
	Qty         int       `gorm:"column:Qty"`
	SellPrice   float64   `gorm:"column:SellPrice"`
	TotalPrice  float64   `gorm:"column:TotalPrice"`
	Remark      string    `gorm:"column:Remark"`
	SKU         string    `gorm:"column:SKU"`
}

type Products struct {
	SKU          string `gorm:"column:SKU"`
	ItemName     string `gorm:"column:ItemName"`
	CurrentStock int    `gorm:"column:CurrentStock"`
}
