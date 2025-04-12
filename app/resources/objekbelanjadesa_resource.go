package resources

import "apidanadesa/app/models"

type ObjekBelanjaDesaResource struct {
	ID           uint   `json:"id"`
	IdKelompok   uint   `json:"id_kelompok"`
	KodeKelompok string `json:"kode_kelompok_belanja"`
	Kelompok     string `json:"kelompok_belanja"`
	IdJenis      uint   `json:"id_jenis"`
	KodeJenis    string `json:"kode_jenis_belanja"`
	Jenis        string `json:"jenis_belanja"`
	KodeObjek    string `json:"kode_objek_belanja"`
	Objek        string `json:"objek_belanja"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at,omitempty"`
}

func newObjekBelanjaDesaResource(m models.ObjekBelanjaDesa) *ObjekBelanjaDesaResource {

	var deletedat string
	if m.DeletedAt.Valid {
		deletedat = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &ObjekBelanjaDesaResource{
		ID:           m.ID,
		IdKelompok:   m.IdKelompok,
		KodeKelompok: m.KelompokBelanjaDesa.Kode,
		Kelompok:     m.KelompokBelanjaDesa.Keterangan,
		IdJenis:      m.IdJenis,
		KodeJenis:    m.JenisBelanjaDesa.Kode,
		Jenis:        m.JenisBelanjaDesa.Keterangan,
		KodeObjek:    m.Kode,
		Objek:        m.Keterangan,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:    deletedat,
	}
}

func GetResponseObjekBelanjaDesa(data []models.ObjekBelanjaDesa) []*ObjekBelanjaDesaResource {
	response := make([]*ObjekBelanjaDesaResource, len(data))
	for index, item := range data {
		response[index] = newObjekBelanjaDesaResource(item)
	}
	return response
}
