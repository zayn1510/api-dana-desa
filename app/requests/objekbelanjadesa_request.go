package requests

import "apidanadesa/app/models"

type ObjekBelanjaDesaRequest struct {
	IdKelompok uint   `json:"id_kelompok" validate:"required"`
	IdJenis    uint   `json:"id_jenis" validate:"required"`
	Kode       string `json:"kode" validate:"required,max=10"`
	Keterangan string `json:"keterangan" validate:"required"`
}

func (req *ObjekBelanjaDesaRequest) ToModelObjekBelanja() *models.ObjekBelanjaDesa {
	return &models.ObjekBelanjaDesa{
		IdKelompok: req.IdKelompok,
		IdJenis:    req.IdJenis,
		Kode:       req.Kode,
		Keterangan: req.Keterangan,
	}
}
