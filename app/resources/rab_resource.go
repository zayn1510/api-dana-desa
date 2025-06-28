package resources

import "apidanadesa/app/models"

type RabResource struct {
	ID            uint    `json:"id"`
	IdKegiatan    uint    `json:"id_kegiatan"`
	PaketKegiatan uint    `json:"paket_kegiatan"`
	Kode          uint    `json:"kode"`
	Anggaran      float64 `json:"anggaran"`
	TahunAnggaran uint    `json:"tahun_anggaran"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	DeletedAt     string  `json:"deleted_at,omitempty"`
}

func NewRabResource(m models.Rab) *RabResource {
	var deleteAt string
	if m.DeletedAt.Valid {
		deleteAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &RabResource{
		ID:            m.ID,
		IdKegiatan:    m.IdKegiatan,
		PaketKegiatan: m.PaketKegiatan,
		Kode:          m.Kode,
		Anggaran:      m.Anggaran,
		TahunAnggaran: m.TahunAnggaran,
		CreatedAt:     m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:     deleteAt,
	}
}
func GetRabResource(data []models.Rab) []*RabResource {
	resources := make([]*RabResource, len(data))
	for i, v := range data {
		resources[i] = NewRabResource(v)
	}
	return resources
}