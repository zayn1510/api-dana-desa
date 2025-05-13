package requests

import "apidanadesa/app/models"
type SumberDanaRequestCreate struct {
	Keterangan string `json:"keterangan" validate:"required,max=100"`
	Kode string `json:"kode" validate:"required,max=10"`
}

type SumberDanaRequestUpdate struct {
	Keterangan string `json:"keterangan,omitempty"`
	Kode string `json:"kode,omitempty"`
}

func (r *SumberDanaRequestCreate) ToModel() *models.SumberDana {
	return &models.SumberDana{
		Keterangan: r.Keterangan,
		Kode: r.Kode,
	}
}