package resources

import "apidanadesa/app/models"

type TahunAnggaranResource struct {
	ID        uint   `json:"id"`
	Tahun     string `json:"tahun"`
	Status    int    `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

func newTahunAnggaranResource(m models.TahunAnggaran) *TahunAnggaranResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &TahunAnggaranResource{
		ID:        m.ID,
		Tahun:     m.Tahun,
		Status:    m.Status,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}

func GetTahunAnggaranResource(data []models.TahunAnggaran) []*TahunAnggaranResource {
	resources := make([]*TahunAnggaranResource, len(data))
	for i := range data {
		resources[i] = newTahunAnggaranResource(data[i])
	}
	return resources
}
