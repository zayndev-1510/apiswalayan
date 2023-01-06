package models

import (
	"apiswayalan.com/rest/api/setup"
	"gorm.io/gorm"
)

var connlogin *gorm.DB

type TblLogin struct {
	Id         int32  `gorm:"primaryKey" json:"id"`
	IdPengguna int32  `json:"id_pengguna"`
	Username   string `gorm:"type:varchar(30)" json:"username"`
	Sandi      string `gorm:"type:varchar(100" json:"sandi"`
	HakAkses   int    `gorm:"type:int(12)" json:"hak_akses"`
	TglBuat    string `json:"tgl_buat"`
}
type InputLogin struct{
	IdPengguna int32  `json:"id_pengguna"`
	Username   string `gorm:"type:varchar(30)" json:"username"`
	Sandi      string `gorm:"type:varchar(100" json:"sandi"`
	HakAkses   int    `gorm:"type:int(12)" json:"hak_akses"`
	TglBuat    string `json:"tgl_buat"`
}

type RequestInputLogin struct{
	InputLogin `json:"akun"`
}

func (TblLogin) TableName() string {
	return "tbl_login"
}

func CreateUser(row *TblLogin) *gorm.DB{

	result:=connlogin.Create(&row)

	return result

}

func UpdateUser(Id int32,row *TblLogin) *gorm.DB{
	stat:=connkasir.Where("id_pengguna=?",Id).Updates(&row)
	return stat
}

func init() {
	connlogin = setup.ConnectDB()
}

func Login(row *TblLogin) *TblLogin {
	var user = TblLogin{}
	connlogin.Find(&user, &row)
	return &user

}
