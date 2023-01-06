package controllers

import (
	"net/http"

	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
	
)


var RegistrasiAkun=func(c *gin.Context){
	
	
	var objek models.RegistrasiModel
	objek.Nama=c.Request.PostFormValue("nama")
	objek.Alamat=c.Request.PostFormValue("alamat")
	objek.NomorHp=c.Request.PostFormValue("nomor_hp")
	objek.Foto=c.Request.PostFormValue("foto")
	objek.Email=c.Request.PostFormValue("email")
	
	result:=models.DaftarAkun(&objek)
	
	
	if result.Error !=nil{ 
		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{"error ":result.Error})
	}else{
		c.AbortWithStatusJSON(http.StatusOK,gin.H{"data ":"Data Suksess"})
	}

}