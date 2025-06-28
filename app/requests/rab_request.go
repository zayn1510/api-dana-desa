package requests

import "apidanadesa/app/models"
	type RabRequest struct {
		IdKegiatan    uint    `json:"id_kegiatan" validate:"required"`
		PaketKegiatan uint    `json:"paket_kegiatan" validate:"required"`
		Kode          uint    `json:"kode" validate:"required"`
		Anggaran      float64 `json:"anggaran" validate:"required"`
		TahunAnggaran uint    `json:"tahun_anggaran" validate:"required"`
	}

	func (r *RabRequest) ToModel() *models.Rab {
		return &models.Rab{
			IdKegiatan:    r.IdKegiatan,
			PaketKegiatan: r.PaketKegiatan,
			Kode:          r.Kode,
			Anggaran:      r.Anggaran,
			TahunAnggaran: r.TahunAnggaran,
		}
	}