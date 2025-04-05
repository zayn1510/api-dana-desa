package resources

import "apidanadesa/app/models"

type KelompokBelanjaDesaResource struct {
	ID           uint   `json:"id"`
	Keterangan   string `json:"keterangan"`
	KodeKelompok string `json:"kode_kelompok_belanja"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at,omitempty"`
}

func newKelompokBelanjaResource(m models.KelompokBelanjaDesa) *KelompokBelanjaDesaResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &KelompokBelanjaDesaResource{
		ID:           m.ID,
		KodeKelompok: m.Kode,
		Keterangan:   m.Keterangan,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:    deletedAt,
	}
}
func GetResponseKelompokBelanja(data []models.KelompokBelanjaDesa) []*KelompokBelanjaDesaResource {
	response := make([]*KelompokBelanjaDesaResource, len(data))
	for i, _ := range response {
		response[i] = newKelompokBelanjaResource(data[i])
	}
	return response
}
