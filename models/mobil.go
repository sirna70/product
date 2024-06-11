package models

type Mobil struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	MerekMobil string `gorm:"varchar(300)" json:"merek_mobil"`
	NamaMobil  string `gorm:"varchar(300)" json:"nama_mobil"`
	StockMobil int64  `gorm:"int" json:"stock_mobil"`
	Deskripsi  string `gorm:"text" json:"deskripsi"`
}
