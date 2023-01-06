package controllers

import (
	"net/http"
	"strconv"

	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
)

type Response struct{
	Count int `json:"count"`
	Data string `json:"data"`
	Message string `json:"message"`
}
var TblBarang models.BarangModels

var GetDataBarang=func(c *gin.Context){
	TblBarang:=models.GetAllData()
	c.IndentedJSON(http.StatusOK,TblBarang)
}
var DetailBarang=func(c *gin.Context)  {
	id,_:=strconv.ParseInt(c.Param("id"),0,0)
	TblBarang,_:=models.GetDataById(id)
	c.IndentedJSON(http.StatusOK,TblBarang)
}
var CreateBarang=func(c *gin.Context){
	errjson:=c.ShouldBindJSON(&TblBarang)
	if(errjson !=nil){
		res := &Response{Count: 0,Data: "",Message: "value not json"}
		c.IndentedJSON(http.StatusBadRequest,res)	
	}else{
		db:=models.CreatedProduk(&TblBarang)
		err:=db.Error
		
		if(err !=nil){
			
			res := &Response{Count: 0,Data: "",Message: "Ada Kesalahan "+err.Error()}
			c.IndentedJSON(http.StatusBadRequest,res)
		
		}else{
			res := &Response{Count: 1,Data: "",Message: "Simpan Data Berhasil"}
			c.IndentedJSON(http.StatusOK,res)
	}
	
	}
}

var UpdateBarang=func(c *gin.Context){
	errjson:=c.ShouldBindJSON(&TblBarang)
	if(errjson !=nil){
		res := &Response{Count: 0,Data: "",Message: "value not json"}
		c.IndentedJSON(http.StatusBadRequest,res)	
	}else{
		db:=models.UpdateProduk(&TblBarang,TblBarang.Id)
		err:=db.Error
		
		if(err !=nil){
			res := &Response{Count: 0,Data: "",Message: "Ada Kesalahan "+err.Error()}
			c.IndentedJSON(http.StatusBadRequest,res)
		
		}else{
			res := &Response{Count: 1,Data: "",Message: "Update Data Berhasil"}
			c.IndentedJSON(http.StatusOK,res)
	}
	
	}
}
var DeleteBarang=func(c *gin.Context){
	id,_:=strconv.ParseInt(c.Param("id"),0,32)
	var obj=models.BarangModels{Id: int32(id)}
	db:=models.DeleteProduk(&obj)
	err:=db.Error
	
	if(err !=nil){
		res := &Response{Count: 0,Data: "",Message: "Ada Kesalahan "+err.Error()}
		c.IndentedJSON(http.StatusBadRequest,res)
	
	}else
		{
			c.AbortWithStatusJSON(http.StatusOK,gin.H{
				"count":1,
				"message":"Hapus Data Berhasil",
				"status":http.StatusOK})
		}
}




