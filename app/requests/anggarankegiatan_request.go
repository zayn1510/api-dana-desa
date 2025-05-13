package requests

import "apidanadesa/app/models"

type AnggaranKegiatanRequestCreate struct {
	IDKegiatan      int     `json:"id_kegiatan" validate:"required"`
	Lokasi          string  `json:"lokasi" validate:"required"`
	Waktu           string  `json:"waktu" validate:"required"`
	IDPerangkatDesa int     `json:"id_perangkat_desa" validate:"required"`
	Keluaran        string  `json:"keluaran" validate:"required"`
	Volume          string  `json:"volume" validate:"required"`
	Pagu            float64 `json:"pagu" validate:"required"`
	IDTahunAnggaran int     `json:"id_tahun_anggaran" validate:"required"`
}

type AnggaranKegiatanRequestUpdate struct {
	IDKegiatan      int     `json:"id_kegiatan,omitempty"`
	Lokasi          string  `json:"lokasi,omitempty"`
	Waktu           string  `json:"waktu,omitempty"`
	IDPerangkatDesa int     `json:"id_perangkat_desa,omitempty"`
	Keluaran        string  `json:"keluaran,omitempty"`
	Volume          string  `json:"volume,omitempty"`
	Pagu            float64 `json:"pagu,omitempty"`
	IDTahunAnggaran   int     `json:"id_tahun_anggaran,omitempty"`
}

func (r *AnggaranKegiatanRequestCreate) ToModel() *models.AnggaranKegiatan {
	return &models.AnggaranKegiatan{
		IDKegiatan:      r.IDKegiatan,
		IDPerangkatDesa: r.IDPerangkatDesa,
		Lokasi:          r.Lokasi,
		Waktu:           r.Waktu,
		Keluaran:        r.Keluaran,
		Volume:          r.Volume,
		Pagu:            r.Pagu,
		IDTahunAnggaran: r.IDTahunAnggaran,
	}
}

func (r *AnggaranKegiatanRequestUpdate) ToModelUpdate() *models.AnggaranKegiatan {
	return &models.AnggaranKegiatan{
		IDKegiatan:      r.IDKegiatan,
		IDPerangkatDesa: r.IDPerangkatDesa,
		Lokasi:          r.Lokasi,
		Waktu:           r.Waktu,
		Keluaran:        r.Keluaran,
		Volume:          r.Volume,
		Pagu:            r.Pagu,
		IDTahunAnggaran: r.IDTahunAnggaran,
	}
}

