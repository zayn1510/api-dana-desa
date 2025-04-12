package requests

import "apidanadesa/app/models"

type ObjekPendapatanRequest struct {
	IdKelompok uint   `json:"id_kelompok" validate:"required"`
	IdJenis    uint   `json:"id_jenis" validate:"required"`
	Kode       string `json:"kode" validate:"required,max=10"`
	Keterangan string `json:"keterangan" validate:"required"`
}

func (req *ObjekPendapatanRequest) ToModelObjekPendapatan() *models.ObjekPendapatan {
	return &models.ObjekPendapatan{
		IdKelompok: req.IdKelompok,
		IdJenis:    req.IdJenis,
		Kode:       req.Kode,
		Keterangan: req.Keterangan,
	}
}
