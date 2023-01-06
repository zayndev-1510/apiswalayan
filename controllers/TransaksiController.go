package controllers

import (
	"net/http"
	"time"

	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type InputTransaksi struct {
	NomorTransaksi string `json:"nomor_transaksi"`
	Status         int    `json:"status"`
}
type requestTransaksi struct {
	Transaksi InputTransaksi `json:"transaksi"`
}

type InputProduk struct {
	Stok   float32 `json:"stok"`
	Id     int32   `json:"id_produk"`
	Jumlah float32 `json:"jumlah"`
}
type RequestProduk struct {
	Produk []InputProduk `json:"produk"`
}

type RequestFilterData struct{
	Startmonth string `json:"startmonth"`
	EndMonth string `json:"endmonth"`
	StartYear string `json:"startyear"`
	EndYear string `json:"endyear"`
}

var tblkeranjang models.TblKeranjang

var GetDataTransaksiAll = func(c *gin.Context) {
	datatransaksi, db := models.GetDataTransaksiAll()
	if db.Error != nil {
		c.JSON(500, gin.H{"message": "Ada Kesalahan pada " + db.Error.Error(), "count": 0})
		return
	}

	if db.RowsAffected == 0 {
		c.JSON(200, gin.H{"message": "Data Kosong ", "count": 0, "data": datatransaksi})
		return
	}
	var datanew []models.TblTranasksi

	for _, f := range datatransaksi {
		t, _ := time.Parse("2006-01-02T15:04:05Z", f.Tgl)
		datanew = append(datanew, models.TblTranasksi{
			Id: f.Id, NomorTransaksi: f.NomorTransaksi,
			NomorKeranjang: f.NomorKeranjang, IdPengguna: f.IdPengguna,
			Tgl: t.Format("2006-01-02"), Total: f.Total, Status: f.Status, RegistrasiModel: f.RegistrasiModel})
	}
	c.IndentedJSON(http.StatusOK,
		gin.H{"count": 1, "data": datanew})
}

var GetDataTransaksiMonth = func(c *gin.Context) {
	var request RequestFilterData
	errjson:=c.ShouldBindJSON(&request)
	if errjson!=nil{
		c.JSON(500,gin.H{"message":"Not Value Json "+errjson.Error(),"count":0})
		return
	}
	datatransaksi, db := models.GetDataTransaksiMonth(request.Startmonth,request.EndMonth)
	if db.Error != nil {
		c.JSON(500, gin.H{"message": "Ada Kesalahan pada " + db.Error.Error(), "count": 0})
		return
	}

	if db.RowsAffected == 0 {
		c.JSON(200, gin.H{"message": "Data Kosong ", "count": 0, "data": datatransaksi})
		return
	}
	var datanew []models.TblTranasksi

	for _, f := range datatransaksi {
		t, _ := time.Parse("2006-01-02T15:04:05Z", f.Tgl)
		datanew = append(datanew, models.TblTranasksi{
			Id: f.Id, NomorTransaksi: f.NomorTransaksi,
			NomorKeranjang: f.NomorKeranjang, IdPengguna: f.IdPengguna,
			Tgl: t.Format("2006-01-02"), Total: f.Total, Status: f.Status, RegistrasiModel: f.RegistrasiModel})
	}
	c.IndentedJSON(http.StatusOK,
		gin.H{"count": 1, "data": datanew})
}

var GetDataTransaksiYear = func(c *gin.Context) {
	var request RequestFilterData
	errjson:=c.ShouldBindJSON(&request)
	if errjson!=nil{
		c.JSON(500,gin.H{"message":"Not Value Json "+errjson.Error(),"count":0})
		return
	}
	datatransaksi, db := models.GetDataTransaksiYear(request.StartYear,request.EndYear)
	if db.Error != nil {
		c.JSON(500, gin.H{"message": "Ada Kesalahan pada " + db.Error.Error(), "count": 0})
		return
	}

	if db.RowsAffected == 0 {
		c.JSON(200, gin.H{"message": "Data Kosong ", "count": 0, "data": datatransaksi})
		return
	}
	var datanew []models.TblTranasksi

	for _, f := range datatransaksi {
		t, _ := time.Parse("2006-01-02T15:04:05Z", f.Tgl)
		datanew = append(datanew, models.TblTranasksi{
			Id: f.Id, NomorTransaksi: f.NomorTransaksi,
			NomorKeranjang: f.NomorKeranjang, IdPengguna: f.IdPengguna,
			Tgl: t.Format("2006-01-02"), Total: f.Total, Status: f.Status, RegistrasiModel: f.RegistrasiModel})
	}
	c.IndentedJSON(http.StatusOK,
		gin.H{"count": 1, "data": datanew})
}
var GetDataTransaksi = func(c *gin.Context) {
	now := time.Now()
	format := now.Format("2006-01-02")
	row := models.TblTranasksi{
		Tgl: format,
	}
	datatransaksi := models.GetDataTransaksi(&row)
	var datanew []models.TblTranasksi

	for _, f := range datatransaksi {
		t, _ := time.Parse("2006-01-02T15:04:05Z", f.Tgl)
		datanew = append(datanew, models.TblTranasksi{
			Id: f.Id, NomorTransaksi: f.NomorTransaksi,
			NomorKeranjang: f.NomorKeranjang, IdPengguna: f.IdPengguna,
			Tgl: t.Format("2006-01-02"), Total: f.Total, Status: f.Status, RegistrasiModel: f.RegistrasiModel})
	}
	c.IndentedJSON(http.StatusOK,
		gin.H{"count": 1, "data": datanew, "tgl": format})

}

var GetDataProdukTransaksi = func(c *gin.Context) {

	row := c.ShouldBindJSON(&tblkeranjang)
	if row != nil {
		c.IndentedJSON(http.StatusOK,
			gin.H{"count": 1, "data": "value not json"})
	} else {
		databarang := models.ProdukTransaksi(&tblkeranjang)
		if len(databarang) == 0 {
			c.IndentedJSON(http.StatusOK,
				gin.H{"count": 0, "data": "not found"})
		} else {
			c.IndentedJSON(http.StatusOK,
				gin.H{"count": 1, "data": databarang})
		}

	}

}

var PerbaruiTransaksi = func(c *gin.Context) {

	var jsontransaksi requestTransaksi
	jsonproduk := RequestProduk{}
	errbarang := c.ShouldBindBodyWith(&jsontransaksi, binding.JSON)
	if errbarang != nil {
		c.IndentedJSON(500, gin.H{"count": 0, "message": "Not value json" + errbarang.Error()})
		return
	}
	errproduk := c.ShouldBindBodyWith(&jsonproduk, binding.JSON)
	if errbarang != nil {
		c.IndentedJSON(500, gin.H{"count": 0, "message": "Not value json" + errproduk.Error()})
		return
	}

	querytransaksiupdate := models.PerbaruiTransaksi(models.TblTranasksi{NomorTransaksi: jsontransaksi.Transaksi.NomorTransaksi},
		map[string]interface{}{"status": jsontransaksi.Transaksi.Status}).Error
	if querytransaksiupdate != nil {
		c.IndentedJSON(500, gin.H{"count": 0, "message": "Perbarui transaksi gagal" + errproduk.Error()})
		return
	}

	for _, data := range jsonproduk.Produk {
		var whereproduk = models.BarangModels{
			Id: data.Id,
		}
		var rowproduk = models.BarangModels{
			Stok: data.Stok - data.Jumlah,
		}
		query := models.UpdateStokBarang(whereproduk, rowproduk).Error
		if query != nil {
			c.IndentedJSON(500, gin.H{"count": 0, "message": "Ada Kesalahan" + query.Error()})
			return
		}
	}
	c.IndentedJSON(200, gin.H{"count": 1, "message": "Berhasil", "data": jsonproduk})
}
