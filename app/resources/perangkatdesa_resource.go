package resources

import (
	"apidanadesa/app/models"
	"fmt"
)

type PerangkatDesaResource struct {
	ID           uint   `json:"id"`
	IDJabatan    uint   `json:"id_jabatan"`
	Jabatan      string `json:"jabatan,omitempty"`
	TempatLahir  string `json:"tempat_lahir"`
	TglLahir     string `json:"tgl_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoSk         string `json:"no_sk"`
	TglSk        string `json:"tgl_sk"`
	NoHandphone  string `json:"no_handphone"`
	Foto         string `json:"foto"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at,omitempty"`
}

func NewPerangkatDesaResource(m models.PerangkatDesa) *PerangkatDesaResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	foto := ""
	if m.Foto != "" {
		foto = fmt.Sprintf("http://localhost:8004/api/v1/preview/%s", m.Foto)
	}

	return &PerangkatDesaResource{ // Mengembalikan pointer agar lebih ringan
		ID:           m.ID,
		IDJabatan:    m.IDJabatan,
		Jabatan:      m.JabatanDesa.Jabatan,
		TempatLahir:  m.TempatLahir,
		TglLahir:     m.TglLahir.Format("2006-01-02 15:04:05"),
		JenisKelamin: m.JenisKelamin,
		NoSk:         m.NoSK,
		TglSk:        m.TglSK.Format("2006-01-02 15:04:05"),
		NoHandphone:  m.NoHandphone,
		Foto:         foto,
		CreatedAt:    m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:    deletedAt,
	}
}

func GetPerangkaDesaResource(data []models.PerangkatDesa) []*PerangkatDesaResource {
	resources := make([]*PerangkatDesaResource, len(data))
	for i, v := range data {
		resources[i] = NewPerangkatDesaResource(v)
	}
	return resources
}
