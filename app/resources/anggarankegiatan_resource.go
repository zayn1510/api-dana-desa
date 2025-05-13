package resources

import (
	"apidanadesa/app/models"
)

type AnggaranKegiatanResource struct {
	ID                uint   `json:"id"`
	IDKegiatan        uint   `json:"id_kegiatan"`
	Kegiatan          string `json:"kegiatan"`
	Lokasi            string `json:"lokasi"`
	Waktu             string `json:"waktu"`
	IDPerangkatDesa   uint   `json:"id_perangkat_desa"`
	NamaPerangkatDesa string `json:"nama_perangkat_desa"`
	Jabatan           string `json:"jabatan"`
	Keluaran          string `json:"keluaran"`
	Volume            string `json:"volume"`
	Pagu              float64 `json:"pagu"`
	IDTahunAnggaran   uint   `json:"id_tahun_anggaran"`
	TahunAnggaran     string `json:"tahun_anggaran"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at"`
}

func NewAnggaranKegiatanResource(m models.AnggaranKegiatan) *AnggaranKegiatanResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &AnggaranKegiatanResource{
		ID:              uint(m.ID),
		IDKegiatan:      uint(m.IDKegiatan),
		Kegiatan: m.Kegiatan.Keterangan,
		Lokasi:           m.Lokasi,
		Waktu:            m.Waktu,
		IDPerangkatDesa:  uint(m.IDPerangkatDesa),
		NamaPerangkatDesa: m.PerangkatDesa.NamaLengkap,
		Jabatan: m.PerangkatDesa.JabatanDesa.Jabatan,
		Keluaran:         m.Keluaran,
		Volume:           m.Volume,
		Pagu:             m.Pagu,
		IDTahunAnggaran: uint(m.IDTahunAnggaran),
		TahunAnggaran:    m.TahunAnggaran.Tahun,
		CreatedAt:        m.CreatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:        deletedAt,
	}
}

func GetAnggaranKegiatanResource(data []models.AnggaranKegiatan) []*AnggaranKegiatanResource {
	resources := make([]*AnggaranKegiatanResource, len(data))
	for i, v := range data {
		resources[i] = NewAnggaranKegiatanResource(v)
	}
	return resources
}

