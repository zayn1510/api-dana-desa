package requests

import "apidanadesa/app/models"

type JabatanDesaRequest struct {
	Jabatan string `json:"jabatan" binding:"required"`
}

func (req *JabatanDesaRequest) ToModelJabatan() *models.JabatanDesa {
	return &models.JabatanDesa{
		Jabatan: req.Jabatan,
	}
}
