package resources

import "apidanadesa/app/models"

type SumberDanaResource struct {
	ID         uint   `json:"id"`
	Keterangan string `json:"keterangan"`
	Kode       string `json:"kode"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}

func NewSumberDana(m models.SumberDana) *SumberDanaResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &SumberDanaResource{
		ID:         m.ID,
		Keterangan: m.Keterangan,
		Kode: m.Kode,
		CreatedAt:  m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:  deletedAt,
	}
}

func GetSumberDanaResource(data []models.SumberDana) []*SumberDanaResource {
	resources := make([]*SumberDanaResource, len(data))
	for i, v := range data {
		resources[i] = NewSumberDana(v)
	}
	return resources
}
