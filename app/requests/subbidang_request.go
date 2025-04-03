package requests

import "apidanadesa/app/models"

type SubBidangRequestCreate struct {
	IdBidang   uint   `json:"id_bidang" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	KodeBidang string `json:"kode_sub_bidang" binding:"required"`
}

func (req *SubBidangRequestCreate) ToModelSubBidang() *models.SubBidang {
	return &models.SubBidang{
		IdBidang:      req.IdBidang,
		Keterangan:    req.Keterangan,
		KodeSubBidang: req.KodeBidang,
	}
}
