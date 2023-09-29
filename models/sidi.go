package models

import (
	"github.com/jinzhu/gorm"
	"github.com/josepnapitupulu/API_Sidi/config"
	)

var db *gorm.DB

type Sidi struct {
	gorm.Model
	NIK string `gorm:""json:"nik"`
	NamaJemaat string `json:"nama_jemaat"`
	TanggalSidi string `json:"tanggal_sidi"`
	TanggalLahir string `json:"tanggal_lahir"`
	NamaPendeta string `json:"nama_pendeta"`
	Alamat string `json:"alamat"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Sidi{})
}

func (b *Sidi) CreateSidi() *Sidi {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllSidis() []Sidi {
	var Sidis []Sidi
	db.Find(&Sidis)
	return Sidis
}

func GetSidibyId(nik int64) (*Sidi, *gorm.DB) {
	var getSidi Sidi
	db := db.Where("NIK=?", nik).Find(&getSidi)
	return &getSidi, db
}

func DeleteSidi(nik int64) Sidi {
	var sidi Sidi
	db.Where("NIK=?", nik).Delete(sidi)
	return sidi 
}