package controllers

import (
	"net/http"
	"strconv"
	"time"

	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
)

var GetDataSuplier = func(c *gin.Context) {

	data, err := models.GetDataSuplier()
	if err.Error != nil {
		c.IndentedJSON(500, gin.H{"message": err.Error.Error(), "count": 0})
		return
	}
	if err.RowsAffected == 0 {
		c.IndentedJSON(200, gin.H{"message": "Data Kosong", "count": 0, "data": data})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Data ada", "count": 1, "data": data})
}

var AddDataSuplier = func(c *gin.Context) {
	var row = models.TblSuplier{}
	now := time.Now()
	row.Tgl = now.Format("2006-01-02")
	errjson := c.ShouldBindJSON(&row)
	if errjson != nil {
		c.IndentedJSON(500, gin.H{"message": errjson.Error(), "count": 0})
		return
	}

	db := models.SaveDataSuplier(&row)
	if db.Error != nil {
		c.IndentedJSON(500, gin.H{"message": db.Error.Error(), "count": 0})
		return
	}
	if db.RowsAffected == 0 {
		c.IndentedJSON(200, gin.H{"message": "Data Gagal Disimpan", "count": 0})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Simpan Data Berhasil", "count": 1})
}

var UpdateSuplier = func(c *gin.Context) {
	var row = models.TblSuplier{}
	errjson := c.ShouldBindJSON(&row)
	if errjson != nil {
		c.IndentedJSON(500, gin.H{"message": errjson.Error(), "count": 0})
		return
	}
	db := models.UpdateDataSuplier(row.Id, &row)
	if db.Error != nil {
		c.IndentedJSON(500, gin.H{"message": db.Error.Error(), "count": 0})
		return
	}
	if db.RowsAffected == 0 {
		c.IndentedJSON(500, gin.H{"message": "Data Gagal Diperbarui", "count": 0})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Perbarui Data Berhasil", "count": 1})
}

var DeleteSuplier = func(c *gin.Context) {

	id, _ := strconv.ParseInt(c.Param("id"), 0, 0)
	var row = models.TblSuplier{
		Id: int32(id),
	}
	db := models.DeleteDataSuplier(&row)
	if db.Error != nil {
		c.IndentedJSON(500, gin.H{"message": db.Error.Error(), "count": 0})
		return
	}
	if db.RowsAffected == 0 {
		c.IndentedJSON(500, gin.H{"message": "Data Gagal Hapus", "count": 0})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Hapus Data Berhasil", "count": 1})
}
