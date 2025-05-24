package resources

import (
	"apidanadesa/app/models"
	"time"
)

type DetailAnggaranKegiatanResource struct {
	ID                 uint64    `json:"id"`
	IDAnggaranKegiatan uint64    `json:"id_anggaran_kegiatan"`
	NamaPaket          string    `json:"nama_paket"`
	Nilai              float64   `json:"nilai"`
	IDPolaKegiatan     uint      `json:"id_pola_kegiatan"`
	Target             string    `json:"target"`
	Uraian             string    `json:"uraian"`
	Satuan             string    `json:"satuan"`
	IDSumberDana       uint      `json:"id_sumber_dana"`
	SifatKegiatan      string    `json:"sifat_kegiatan"`
	LokasiKegiatan     string    `json:"lokasi_kegiatan"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
func NewDetailAnggaranKegiatanResource(m models.DetailAnggaranKegiatan) DetailAnggaranKegiatanResource {
	return DetailAnggaranKegiatanResource{
		ID:                 m.ID,
		IDAnggaranKegiatan: m.IDAnggaranKegiatan,
		NamaPaket:          m.NamaPaket,
		Nilai:              m.Nilai,
		IDPolaKegiatan:     m.IDPolaKegiatan,
		Target:             m.Target,
		Uraian:             m.Uraian,
		Satuan:             m.Satuan,
		IDSumberDana:       m.IDSumberDana,
		SifatKegiatan:      m.SifatKegiatan,
		LokasiKegiatan:     m.LokasiKegiatan,
		CreatedAt:          m.CreatedAt,
		UpdatedAt:          m.UpdatedAt,
	}
}

func NewDetailAnggaranKegiatanResourceList(data []models.DetailAnggaranKegiatan) []DetailAnggaranKegiatanResource {
	var results []DetailAnggaranKegiatanResource
	for _, item := range data {
		results = append(results, NewDetailAnggaranKegiatanResource(item))
	}
	return results
}
