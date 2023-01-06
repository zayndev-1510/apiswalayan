package models

import (
	"apiswayalan.com/rest/api/setup"
	"gorm.io/gorm"
)

var connkasir *gorm.DB

type TblKasir struct {
	Id           int32  `json:"id"`
	NamaKasir    string `json:"nama_kasir"`
	NomorHp      string `json:"nomor_hp"`
	Alamat       string `json:"alamat"`
	JenisKelamin string `json:"jenis_kelamin"`
	TglLahir     string `json:"tgl_lahir"`
	Pendidikan   string `json:"pendidikan"`
	Foto         string `json:"foto"`
	TglBuat      string `json:"tgl_buat"`
}
type DataRelasi struct {
	IdKasir      int32  `json:"id_kasir"`
	NamaKasir    string `json:"nama_kasir"`
	NomorHp      string `json:"nomor_hp"`
	Alamat       string `json:"alamat"`
	JenisKelamin string `json:"jenis_kelamin"`
	TglLahir     string `json:"tgl_lahir"`
	Pendidikan   string `json:"pendidikan"`
	Foto         string `json:"foto"`
	TglBuat      string `json:"tgl_buat"`
	Username     string `json:"username"`
	Sandi        string `json:"sandi"`
}

func (TblKasir) TableName() string {
	return "tbl_kasir"
}
func init() {
	connkasir = setup.ConnectDB()
}

func DataKasir() ([]DataRelasi, *gorm.DB) {
	var result []DataRelasi
	x := conntransaksi.Table("tbl_login as login").
		Joins("JOIN tbl_kasir ON tbl_kasir.id=login.id_pengguna").
		Select("tbl_kasir.id as id_kasir", "tbl_kasir.nama_kasir", "tbl_kasir.nomor_hp",
			"tbl_kasir.alamat", "tbl_kasir.jenis_kelamin", "tbl_kasir.tgl_lahir", "tbl_kasir.pendidikan",
			"tbl_kasir.foto", "tbl_kasir.tgl_buat", "login.username", "login.sandi").
		Find(&result)
	return result, x
}

func CreateKasir(row *TblKasir) *gorm.DB {
	stat := connkasir.Create(&row)
	return stat
}

func UpdateKasir(id int32, row *TblKasir) *gorm.DB {
	stat := connkasir.Where("id=?", id).Updates(&row)
	return stat
}

func DeleteKasir(row *TblKasir) *gorm.DB {
	stat := connkasir.Delete(&row)
	var rowlogin = TblLogin{
		Id: row.Id,
	}
	connkasir.Delete(&rowlogin)
	return stat
}
