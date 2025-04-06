package requests

import "apidanadesa/app/models"

type SubBidangRequestCreate struct {
	IdBidang   uint   `json:"id_bidang" validate:"required"`
	Keterangan string `json:"keterangan" validate:"required"`
	KodeBidang string `json:"kode_sub_bidang" validate:"required"`
}

func (req *SubBidangRequestCreate) ToModelSubBidang() *models.SubBidang {
	return &models.SubBidang{
		IdBidang:      req.IdBidang,
		Keterangan:    req.Keterangan,
		KodeSubBidang: req.KodeBidang,
	}
}
func (req *SubBidangRequestCreate) ToModelSubBidangCsv() models.SubBidang {
	return models.SubBidang{
		IdBidang:      req.IdBidang,
		Keterangan:    req.Keterangan,
		KodeSubBidang: req.KodeBidang,
	}
}
