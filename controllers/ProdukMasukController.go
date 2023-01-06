package controllers

import (
	"time"

	"apiswayalan.com/rest/api/models"
	"github.com/gin-gonic/gin"
)

type RequestJson struct {
	ProdukMasuk  []models.ProdukMasukModel `json:"produkmasuk"`
	ProdukUpdate []models.BarangModels     `json:"produkupdate"`
}

type RequestFilterProdukMasuk struct {
	Startmonth string `json:"startmonth"`
	Endmonth   string `json:"endmonth"`
	Startyear  string `json:"startyear"`
	Endyear    string `json:"endyear"`
	Aksi       int32  `json:"aksi"`
}

var GetDataProdukMasuk = func(c *gin.Context) {
	tgl := c.Param("tgl")
	result, err := models.GetDataProdukMasuk(tgl)

	if err.Error != nil {

		c.JSON(500, map[string]interface{}{"count": 0, "message": "Ada Kesalahan di " + err.Error.Error()})

		return
	}
	if err.RowsAffected == 0 {
		c.JSON(200, map[string]interface{}{"count": 0, "message": "Data Kosong", "data": result, "tgl": tgl})
		return
	}
	var datanew []models.ProdukMasukModelRelasi

	var totalproduk = 0
	var totalstoklama = 0
	var totalstokbaru = 0
	var totalhargaproduk = 0
	for _, f := range result {
		t, _ := time.Parse("2006-01-02T15:04:05Z", f.Tgl)
		totalproduk = totalproduk + int(f.Jumlah)
		totalstoklama = totalstoklama + int(f.StokLama)
		totalstokbaru = totalstokbaru + (int(f.StokLama) + int(f.Jumlah))
		totalhargaproduk += int(f.Harga) * int(f.Jumlah)
		datanew = append(datanew, models.ProdukMasukModelRelasi{
			Id: f.Id, NamaBarang: f.NamaBarang, IdBarang: f.IdBarang, IdSuplier: f.IdSuplier, Harga: f.Harga,
			Tgl: t.Format("2006-01-02"), Jumlah: f.Jumlah, StokLama: f.StokLama, NamaSuplier: f.NamaSuplier})
	}
	c.JSON(200, map[string]interface{}{"count": 1, "message": "Data Found", "data": datanew,
		"totalproduk": totalproduk, "stoklama": totalstoklama, "stokbaru": totalstokbaru, "totalhargaproduk": totalhargaproduk})
}
var GetDataProdukMasukFilter = func(c *gin.Context) {
	var json = RequestFilterProdukMasuk{}
	var datanew []models.ProdukMasukModelRelasi
	var totalproduk = 0
	var totalstoklama = 0
	var totalstokbaru = 0
	var totalhargaproduk = 0

	errjson := c.ShouldBindJSON(&json)
	if errjson != nil {
		c.JSON(500, map[string]interface{}{"count": 0, "message": "This Not Value Json " + errjson.Error()})
	}
	if json.Aksi == 1 {
		result, err := models.GetDataProdukMasukFilterMonth(json.Startmonth, json.Endmonth)
		if err.Error != nil {

			c.JSON(500, map[string]interface{}{"count": 0, "message": "Ada Kesalahan di " + err.Error.Error()})
			return
		}
		if err.RowsAffected == 0 {
			c.JSON(200, map[string]interface{}{"count": 0, "message": "Data Kosong", "data": result})
			return
		}
		for _, f := range result {
			t, _ := time.Parse("2006-01-02T15:04:05Z", f.Tgl)
			totalproduk = totalproduk + int(f.Jumlah)
			totalstoklama = totalstoklama + int(f.StokLama)
			totalstokbaru = totalstokbaru + (int(f.StokLama) + int(f.Jumlah))
			totalhargaproduk += int(f.Harga) * int(f.Jumlah)
			datanew = append(datanew, models.ProdukMasukModelRelasi{
				Id: f.Id, NamaBarang: f.NamaBarang, IdBarang: f.IdBarang, IdSuplier: f.IdSuplier, Harga: f.Harga,
				Tgl: t.Format("2006-01-02"), Jumlah: f.Jumlah, StokLama: f.StokLama, NamaSuplier: f.NamaSuplier})
		}
	} else {
		result, err := models.GetDataProdukMasukFilterYear(json.Startyear, json.Endyear)
		if err.Error != nil {

			c.JSON(500, map[string]interface{}{"count": 0, "message": "Ada Kesalahan di " + err.Error.Error()})

			return
		}
		if err.RowsAffected == 0 {
			c.JSON(200, map[string]interface{}{"count": 0, "message": "Data Kosong", "data": result})
			return
		}
		for _, f := range result {
			t, _ := time.Parse("2006-01-02T15:04:05Z", f.Tgl)
			totalproduk = totalproduk + int(f.Jumlah)
			totalstoklama = totalstoklama + int(f.StokLama)
			totalstokbaru = totalstokbaru + (int(f.StokLama) + int(f.Jumlah))
			totalhargaproduk += int(f.Harga) * int(f.Jumlah)
			datanew = append(datanew, models.ProdukMasukModelRelasi{
				Id: f.Id, NamaBarang: f.NamaBarang, IdBarang: f.IdBarang, IdSuplier: f.IdSuplier, Harga: f.Harga,
				Tgl: t.Format("2006-01-02"), Jumlah: f.Jumlah, StokLama: f.StokLama, NamaSuplier: f.NamaSuplier})
		}
	}

	c.JSON(200, map[string]interface{}{"count": 1, "message": "Data Found", "data": datanew,
		"totalproduk": totalproduk, "stoklama": totalstoklama, "stokbaru": totalstokbaru, "totalhargaproduk": totalhargaproduk})
}

var SaveDataProdukMasuk = func(c *gin.Context) {
	var json = RequestJson{}
	errjson := c.ShouldBindJSON(&json)

	if errjson != nil {
		c.JSON(500, map[string]interface{}{"count": 0, "message": "Ada Kesalahan di " + errjson.Error()})
		return
	}

	db := models.SaveDataProdukMasuk(json.ProdukMasuk)
	errdb := db.Error
	result := db.RowsAffected
	if errdb != nil {
		c.JSON(500, map[string]interface{}{"count": 0, "message": "Ada Kesalahan di " + errdb.Error()})
		return
	}
	if result == 0 {
		c.JSON(500, map[string]interface{}{"count": 0, "message": "Simpan data produk masuk gagal"})
		return
	}

	for _, data := range json.ProdukUpdate {
		models.UpdateStokBaragArray(data.Id, data)
	}

	c.JSON(200, map[string]interface{}{"message": "Perbarui data stok barang berhasil", "count": 1})

}
