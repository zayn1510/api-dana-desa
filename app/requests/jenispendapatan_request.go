package requests

import "apidanadesa/app/models"

type JenisPendapatanRequest struct {
	IdKelompok uint   `json:"id_kelompok" validate:"required"`
	Kode       string `json:"kode" validate:"required,max=10"`
	Keterangan string `json:"keterangan" validate:"required"`
}

func (req *JenisPendapatanRequest) ToModelJenisPendapatan() *models.JenisPendapatan {
	return &models.JenisPendapatan{
		IdKelompok: req.IdKelompok,
		Kode:       req.Kode,
		Keterangan: req.Keterangan,
	}
}
