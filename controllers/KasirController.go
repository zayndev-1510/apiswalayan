package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
)

type RequestInputDataKasir struct {
	models.TblLogin `json:"akun"`
	models.TblKasir `json:"kasir"`
}

var Datakasir = func(c *gin.Context) {

	datakasir, err := models.DataKasir()
	if err.Error != nil {
		c.IndentedJSON(500, gin.H{"message": "Ada Kesalahan " + err.Error.Error(), "count": 0})
		return

	}
	var datanew []models.DataRelasi
	for _, data := range datakasir {
		tgl_lahir, _ := time.Parse("2006-01-02T15:04:05Z", data.TglLahir)
		tgl_buat, _ := time.Parse("2006-01-02T15:04:05Z", data.TglBuat)
		datanew = append(datanew, models.DataRelasi{
			IdKasir: data.IdKasir, NamaKasir: data.NamaKasir, Alamat: data.Alamat, Pendidikan: data.Pendidikan,
			JenisKelamin: data.JenisKelamin, TglLahir: tgl_lahir.Format("2006-01-02"), NomorHp: data.NomorHp,
			TglBuat: tgl_buat.Format("2006-01-02"), Foto: data.Foto, Username: data.Username, Sandi: data.Sandi})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": datanew, "count": 1})

}

var CreateKasir = func(c *gin.Context) {

	jsondata := map[string]interface{}{}

	err := c.ShouldBindJSON(&jsondata)
	if err != nil {
		c.IndentedJSON(500, gin.H{"message": "Ada Kesalahan " + err.Error(), "count": 0})
		return
	}
	jsonkasir, _ := json.Marshal(jsondata)
	inputkasir := RequestInputDataKasir{}
	json.Unmarshal(jsonkasir, &inputkasir)
	var now = time.Now()
	inputkasir.TblKasir.TglBuat = now.Format("2006-01-02")
	inputkasir.TblLogin.TglBuat = now.Format("2006-01-02")
	stat := models.CreateKasir(&inputkasir.TblKasir)
	if stat.Error != nil {
		c.IndentedJSON(500, gin.H{"message": "Ada Kesalahan " + stat.Error.Error(), "count": 0})
		return
	}
	inputkasir.TblLogin.IdPengguna = inputkasir.TblKasir.Id
	querylogin := models.CreateUser(&inputkasir.TblLogin)
	if querylogin.Error != nil {
		c.IndentedJSON(500, gin.H{"message": "Ada Kesalahan " + querylogin.Error.Error(), "count": 0})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Buat Data Kasir Berhasil", "count": 1})

}

var UpdateKasir = func(c *gin.Context) {
	jsondata := map[string]interface{}{}

	err := c.ShouldBindJSON(&jsondata)
	if err != nil {
		c.JSON(500, gin.H{"message": "Ada Kesalahan " + err.Error(), "count": 0})
	}
	jsonkasir, _ := json.Marshal(jsondata)
	inputkasir := RequestInputDataKasir{}
	json.Unmarshal(jsonkasir, &inputkasir)
	updatekasir := models.UpdateKasir(inputkasir.TblKasir.Id, &inputkasir.TblKasir)
	if updatekasir.Error != nil {
		c.IndentedJSON(500, gin.H{"message": "Ada Kesalahan " + updatekasir.Error.Error(), "count": 0})
		return
	}
	Id := inputkasir.TblKasir.Id
	queryupdateuser := models.UpdateUser(Id, &inputkasir.TblLogin)
	if queryupdateuser.Error != nil {
		c.IndentedJSON(500, gin.H{"message": "Ada Kesalahan " + queryupdateuser.Error.Error(), "count": 0})
		return
	}
	c.JSON(200, map[string]interface{}{"message": "Data Berhasil Di Update", "count": 1})
}

var DeleteKasir = func(c *gin.Context) {
	var kasir models.TblKasir
	err := c.ShouldBindJSON(&kasir)
	if err != nil {
		c.JSON(500, gin.H{"message": "Value Bukan Json" + err.Error(), "count": 0})
		return
	}
	StatHapusKasir := models.DeleteKasir(&kasir)
	if StatHapusKasir.Error != nil {
		c.JSON(500, gin.H{"message": "Ada Kesalahan " + StatHapusKasir.Error.Error(), "count": 0})
		return
	}

	if StatHapusKasir.RowsAffected == 0 {
		c.JSON(500, gin.H{"message": "Hapus Data Gagal ", "count": 0})
		return
	}

	c.JSON(200, map[string]interface{}{"message": "Data Berhasil Di Hapus", "count": 1})

}


var Coba=func(c *gin.Context){
	var request=RequestInputDataKasir{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, gin.H{"message": "Ada Kesalahan " + err.Error(), "count": 0})
	}
	c.JSON(200, map[string]interface{}{"message": "Data Berhasil Di Hapus", "count": 1,"data":request.TblKasir})
} 