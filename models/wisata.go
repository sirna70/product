package models

type Wisata struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NamaDaerah   string `gorm:"varchar(300)" json:"nama_daerah"`
	TempatWisata string `gorm:"text" json:"tempat_wisata"`
	TarifWisata  int64  `gorm:"int" json:"tarif_wisata"`
}
