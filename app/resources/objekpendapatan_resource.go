package resources

import "apidanadesa/app/models"

type ObjekPendapatanResource struct {
	ID           uint   `json:"id"`
	IdKelompok   uint   `json:"id_kelompok"`
	KodeKelompok string `json:"kode_kelompok_pendapatan"`
	Kelompok     string `json:"kelompok_pendapatan"`
	IdJenis      uint   `json:"id_jenis"`
	KodeJenis    string `json:"kode_jenis_pendapatan"`
	Jenis        string `json:"jenis_pendapatan"`
	KodeObjek    string `json:"kode_objek_pendapatan"`
	Objek        string `json:"objek_pendapatan"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at,omitempty"`
}

func newObjekPendapatanResource(m models.ObjekPendapatan) *ObjekPendapatanResource {

	var deletedat string
	if m.DeletedAt.Valid {
		deletedat = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &ObjekPendapatanResource{
		ID:           m.ID,
		IdKelompok:   m.IdKelompok,
		KodeKelompok: m.KelompokPendapatan.Kode,
		Kelompok:     m.KelompokPendapatan.Keterangan,
		IdJenis:      m.IdJenis,
		KodeJenis:    m.JenisPendapatan.Kode,
		Jenis:        m.JenisPendapatan.Keterangan,
		KodeObjek:    m.Kode,
		Objek:        m.Keterangan,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:    deletedat,
	}
}

func GetResponseObjekPendapatan(data []models.ObjekPendapatan) []*ObjekPendapatanResource {
	response := make([]*ObjekPendapatanResource, len(data))
	for index, item := range data {
		response[index] = newObjekPendapatanResource(item)
	}
	return response
}
