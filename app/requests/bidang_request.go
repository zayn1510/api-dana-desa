package requests

import "apidanadesa/app/models"

type BidangRequestCreate struct {
	Keterangan string `json:"keterangan" validate:"required,max=100"`
	KodeBidang string `json:"kode_bidang" validate:"required,max=10,numeric"`
}

func (r *BidangRequestCreate) ToModel() *models.Bidang {
	return &models.Bidang{
		Keterangan: r.Keterangan,
		KodeBidang: r.KodeBidang,
	}
}
func (r *BidangRequestCreate) ToModelCsv() models.Bidang {
	return models.Bidang{
		Keterangan: r.Keterangan,
		KodeBidang: r.KodeBidang,
	}
}
