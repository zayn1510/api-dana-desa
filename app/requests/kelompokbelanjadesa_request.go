package requests

import (
	"apidanadesa/app/models"
)

type KelompokBelanjaDesaRequest struct {
	Kode       string `validate:"required,min=1,max=10" json:"kode"`
	Keterangan string `validate:"required" json:"keterangan"`
}

func (req *KelompokBelanjaDesaRequest) ToModelKelompokBelanja() *models.KelompokBelanjaDesa {
	return &models.KelompokBelanjaDesa{
		Kode:       req.Kode,
		Keterangan: req.Keterangan,
	}
}
