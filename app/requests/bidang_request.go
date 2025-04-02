package requests

import "apidanadesa/app/models"
type BidangRequestCreate struct {
	Keterangan string `json:"keterangan" binding:"required"`
	KodeBidang string `json:"kode_bidang" binding:"required"`
}

func (r *BidangRequestCreate) ToModel() *models.Bidang {
	return&models.Bidang{
		Keterangan: r.Keterangan,
		KodeBidang: r.KodeBidang,
	}
}
