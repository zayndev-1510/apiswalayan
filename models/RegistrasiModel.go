package models

import (
	"apiswayalan.com/rest/api/setup"
	"gorm.io/gorm"
)
var TblRegistrasi *gorm.DB
type RegistrasiModel struct{
	IdRegister int32 	`gorm:"primaryKey" json:"id_register"`
	Email string 		`gorm:"type:varchar(30)" json:"email"`
	Nama string 		`gorm:"type:varchar(50)" json:"nama"`
	NomorHp string  	`gorm:"type:varchar(12)" json:"nomor_hp"`
	Alamat string 		`gorm:"type:varchar(30)" json:"alamat"`
	Foto string			`gorm:"type:varchar(50)" json:"foto"`
}

func (RegistrasiModel) TableName() string {
    return "tbl_registrasi"
}

func init(){
	TblRegistrasi=setup.ConnectDB()
}

func DaftarAkun(row *RegistrasiModel) *gorm.DB{
	result:=TblRegistrasi.Create(&row)
	return result
	
}