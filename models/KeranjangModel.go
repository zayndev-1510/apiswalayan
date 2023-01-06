package models


type TblKeranjang struct {
	IdKeranjajng              int32           `json:"id_keranjang"`
	IdProduk int32 `json:"id_produk"`
	NomorKeranjang  string          `json:"nomor_keranjang"`
	IdBarang  		int32          	`json:"id_barang"`
	NamaBarang 		string			`json:"nama_barang"`
	Jumlah  		float32         `json:"jumlah"`
	Harga          	float32         `json:"harga"`
   	Status          int32           `json:"status"`
	Stok 			float32			`json:"stok"`
}

func (TblKeranjang) TableName() string {
	return "tbl_keranjang"
}

