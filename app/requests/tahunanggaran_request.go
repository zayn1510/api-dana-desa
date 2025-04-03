package requests

import "apidanadesa/app/models"

type TahunAnggaranRequest struct {
	Tahun  string `json:"tahun" binding:"required,max=30"`
	Status int    `json:"status" binding:"max=1"`
}

func (t *TahunAnggaranRequest) ToModelTahunAnggaran() *models.TahunAnggaran {
	return &models.TahunAnggaran{
		Tahun:  t.Tahun,
		Status: t.Status,
	}
}
