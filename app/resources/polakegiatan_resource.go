package resources

import "apidanadesa/app/models"

type PolaKegiatanResource struct {
	ID         uint   `json:"id"`
	Keterangan string `json:"keterangan"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}

func NewPolaKegiatanResource(m models.PolaKegiatan) *PolaKegiatanResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &PolaKegiatanResource{
		ID:         m.ID,
		Keterangan: m.Keterangan,
		CreatedAt:  m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:  deletedAt,
	}
}

func GetPolaKegiatanResource(data []models.PolaKegiatan) []*PolaKegiatanResource {
	resources := make([]*PolaKegiatanResource, len(data))
	for i, v := range data {
		resources[i] = NewPolaKegiatanResource(v)
	}
	return resources
}
