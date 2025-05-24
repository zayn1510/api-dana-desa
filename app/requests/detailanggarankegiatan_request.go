package requests

import "apidanadesa/app/models"

type CreateDetailAnggaranKegiatanRequest struct {
	IDAnggaranKegiatan uint64  `json:"id_anggaran_kegiatan" validate:"required"`
	NamaPaket          string  `json:"nama_paket" validate:"required,max=100"`
	Nilai              float64 `json:"nilai" validate:"required"`
	IDPolaKegiatan     uint    `json:"id_pola_kegiatan" validate:"required"`
	Target             string  `json:"target" validate:"required,max=200"`
	Uraian             string  `json:"uraian" validate:"required,max=200"`
	Satuan             string  `json:"satuan" validate:"required,max=200"`
	IDSumberDana       uint    `json:"id_sumber_dana" validate:"required"`
	SifatKegiatan      string  `json:"sifat_kegiatan" validate:"required,max=200"`
	LokasiKegiatan     string  `json:"lokasi_kegiatan" validate:"required,max=200"`
}

func (r *CreateDetailAnggaranKegiatanRequest) ToModel() *models.DetailAnggaranKegiatan {
	return &models.DetailAnggaranKegiatan{
		IDAnggaranKegiatan: r.IDAnggaranKegiatan,
		NamaPaket:          r.NamaPaket,
		Nilai:              r.Nilai,
		IDPolaKegiatan:     r.IDPolaKegiatan,
		Target:             r.Target,
		Uraian:             r.Uraian,
		Satuan:             r.Satuan,
		IDSumberDana:       r.IDSumberDana,
		SifatKegiatan:      r.SifatKegiatan,
		LokasiKegiatan:     r.LokasiKegiatan,
	}
}
