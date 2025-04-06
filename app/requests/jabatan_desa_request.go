package requests

import "apidanadesa/app/models"

type JabatanDesaRequest struct {
	Jabatan string `json:"jabatan" validate:"required"`
}

func (req *JabatanDesaRequest) ToModelJabatan() *models.JabatanDesa {
	return &models.JabatanDesa{
		Jabatan: req.Jabatan,
	}
}
func (req *JabatanDesaRequest) ToModelJabatanCsv() models.JabatanDesa {
	return models.JabatanDesa{
		Jabatan: req.Jabatan,
	}
}
