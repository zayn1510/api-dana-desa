package requests

import "apidanadesa/app/models"

type JenisBelanjaDesaRequest struct {
	IdKelompok uint   `json:"id_kelompok" validate:"required"`
	Kode       string `json:"kode" validate:"required,max=10"`
	Keterangan string `json:"keterangan" validate:"required"`
}

func (req *JenisBelanjaDesaRequest) ToModelJenisBelanja() *models.JenisBelanjaDesa {
	return &models.JenisBelanjaDesa{
		IdKelompok: req.IdKelompok,
		Kode:       req.Kode,
		Keterangan: req.Keterangan,
	}
}
