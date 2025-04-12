package resources

import "apidanadesa/app/models"

type JenisBelanjaDesaResource struct {
	ID           uint   `json:"id"`
	IdKelompok   uint   `json:"id_kelompok"`
	KodeKelompok string `json:"kode_kelompok_belanja"`
	Kelompok     string `json:"kelompok_belanja"`
	KodeJenis    string `json:"kode_jenis_belanja"`
	Keterangan   string `json:"jenis_belanja"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at,omitempty"`
}

func newJenisBelanjaDesaResource(m models.JenisBelanjaDesa) *JenisBelanjaDesaResource {

	var deletedat string
	if m.DeletedAt.Valid {
		deletedat = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &JenisBelanjaDesaResource{
		ID:           m.ID,
		IdKelompok:   m.IdKelompok,
		KodeKelompok: m.KelompokBelanjaDesa.Kode,
		Kelompok:     m.KelompokBelanjaDesa.Keterangan,
		KodeJenis:    m.Kode,
		Keterangan:   m.Keterangan,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:    deletedat,
	}
}

func GetResponseJenisBelanjaDesa(data []models.JenisBelanjaDesa) []*JenisBelanjaDesaResource {
	response := make([]*JenisBelanjaDesaResource, len(data))
	for index, item := range data {
		response[index] = newJenisBelanjaDesaResource(item)
	}
	return response
}
