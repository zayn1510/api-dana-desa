package resources

import "apidanadesa/app/models"

type JenisPendapatanResource struct {
	ID           uint   `json:"id"`
	IdKelompok   uint   `json:"id_kelompok"`
	KodeKelompok string `json:"kode_kelompok_pendapatan"`
	Kelompok     string `json:"kelompok_pendapatan"`
	KodeJenis    string `json:"kode_jenis_pendapatan"`
	Keterangan   string `json:"jenis_pendapatan"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at,omitempty"`
}

func newJenisPendapatanResource(m models.JenisPendapatan) *JenisPendapatanResource {

	var deletedat string
	if m.DeletedAt.Valid {
		deletedat = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &JenisPendapatanResource{
		ID:           m.ID,
		IdKelompok:   m.IdKelompok,
		KodeKelompok: m.KelompokPendapatan.Kode,
		Kelompok:     m.KelompokPendapatan.Keterangan,
		KodeJenis:    m.Kode,
		Keterangan:   m.Keterangan,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:    deletedat,
	}
}

func GetJenisPendapatanResource(data []models.JenisPendapatan) []*JenisPendapatanResource {
	response := make([]*JenisPendapatanResource, len(data))
	for index, item := range data {
		response[index] = newJenisPendapatanResource(item)
	}
	return response
}
