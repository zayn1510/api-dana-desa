package resources

import "apidanadesa/app/models"

type JabatanDesaResource struct {
	ID        uint   `json:"id"`
	Jabatan   string `json:"jabatan"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

func newJabatanDesaResource(m models.JabatanDesa) *JabatanDesaResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &JabatanDesaResource{
		ID:        m.ID,
		Jabatan:   m.Jabatan,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: deletedAt,
	}
}
func GetResponseJabatanDesa(data []models.JabatanDesa) []*JabatanDesaResource {
	response := make([]*JabatanDesaResource, len(data))
	for i, _ := range response {
		response[i] = newJabatanDesaResource(data[i])
	}
	return response
}
