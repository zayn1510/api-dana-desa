package resources

import "apidanadesa/app/models"

type KegiatanResource struct {
	ID            uint   `json:"id"`
	IdBidang      uint   `json:"id_bidang"`
	KodeBidang    string `json:"kode_bidang"`
	Bidang        string `json:"bidang"`
	IdSubBidang   uint   `json:"id_sub_bidang"`
	KodeSubBidang string `json:"kode_sub_bidang"`
	SubBidang     string `json:"sub_bidang"`
	Kegiatan      string `json:"kegiatan,omitempty"`
	KodeKegiatan  string `json:"kode_kegiatan,omitempty"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at,omitempty"`
}

func newKegiatanResource(m models.Kegiatan) *KegiatanResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &KegiatanResource{
		ID:            m.ID,
		IdBidang:      m.IdBidang,
		IdSubBidang:   m.IdSubBidang,
		KodeBidang:    m.Bidang.KodeBidang,
		Bidang:        m.Bidang.Keterangan,
		KodeSubBidang: m.SubBidang.KodeSubBidang,
		SubBidang:     m.SubBidang.Keterangan,
		KodeKegiatan:  m.KodeKegiatan,
		Kegiatan:      m.Keterangan,
		CreatedAt:     m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:     deletedAt,
	}
}

func GetKegiatanResponse(data []models.Kegiatan) []*KegiatanResource {
	resources := make([]*KegiatanResource, len(data))
	for i, v := range data {
		resources[i] = newKegiatanResource(v)
	}
	return resources
}
