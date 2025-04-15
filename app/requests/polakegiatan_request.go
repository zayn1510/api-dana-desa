package requests

import "apidanadesa/app/models"

type PolaKegiatanCreate struct {
	Keterangan string `json:"keterangan" validate:"required,max=100"`
}

func (r *PolaKegiatanCreate) ToModelPolaKegiatan() *models.PolaKegiatan {
	return &models.PolaKegiatan{
		Keterangan: r.Keterangan,
	}
}
