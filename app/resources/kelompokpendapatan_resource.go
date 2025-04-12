package resources

import "apidanadesa/app/models"

type KelompokPendapatanResource struct {
	ID         uint   `json:"id"`
	Kode       string `json:"kode_kelompok_pendapatan"`
	Keterangan string `json:"keterangan_pendapatan"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}
type KelompokPendapatanCSVRow struct {
	Keterangan string `json:"keterangan"`
	Kode       string `json:"kode_kelompok_pendapatan"`
	Row        int    `json:"row"`
	Error      string `json:"error,omitempty"`
}

func newKelompokPendapatanResource(m models.KelompokPendapatan) *KelompokPendapatanResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &KelompokPendapatanResource{
		ID:         m.ID,
		Kode:       m.Kode,
		Keterangan: m.Keterangan,
		CreatedAt:  m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:  deletedAt,
	}
}

func GetKelompokPendapatnResource(data []models.KelompokPendapatan) []*KelompokPendapatanResource {
	responses := make([]*KelompokPendapatanResource, len(data))
	for i, value := range data {
		responses[i] = newKelompokPendapatanResource(value)
	}
	return responses
}
