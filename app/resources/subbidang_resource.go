package resources

import "apidanadesa/app/models"

type SubBidangResource struct {
	ID            uint   `json:"id"`
	KodeBidang    string `json:"kode_bidang"`
	Bidang        string `json:"bidang"`
	Keterangan    string `json:"keterangan"`
	KodeSubBidang string `json:"kode_sub_bidang"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at,omitempty"`
}

func newSubBidangResource(m models.SubBidang) *SubBidangResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}
	return &SubBidangResource{
		ID:            m.ID,
		KodeBidang:    m.Bidang.KodeBidang,
		Bidang:        m.Bidang.Keterangan,
		Keterangan:    m.Keterangan,
		KodeSubBidang: m.KodeSubBidang,
		CreatedAt:     m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:     deletedAt,
	}
}

func GetSubBidangResponse(data []models.SubBidang) []*SubBidangResource {
	resources := make([]*SubBidangResource, len(data))
	for i, v := range data {
		resources[i] = newSubBidangResource(v)
	}
	return resources
}
