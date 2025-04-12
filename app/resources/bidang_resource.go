package resources

import "apidanadesa/app/models"

type BidangResource struct {
	ID         uint   `json:"id"`
	Keterangan string `json:"keterangan"`
	KodeBidang string `json:"kode_bidang"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}
type BidangCSVRow struct {
	Keterangan string `json:"keterangan"`
	KodeBidang string `json:"kode_bidang"`
	Row        int    `json:"row"`
	Error      string `json:"error,omitempty"`
}

func NewBidangResource(m models.Bidang) *BidangResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &BidangResource{ // Mengembalikan pointer agar lebih ringan
		ID:         m.ID,
		Keterangan: m.Keterangan,
		KodeBidang: m.KodeBidang,
		CreatedAt:  m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:  m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:  deletedAt,
	}
}

func GetBidangResource(data []models.Bidang) []*BidangResource {
	resources := make([]*BidangResource, len(data))
	for i, v := range data {
		resources[i] = NewBidangResource(v)
	}
	return resources
}
