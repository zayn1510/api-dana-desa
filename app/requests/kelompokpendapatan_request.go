package requests

import "apidanadesa/app/models"

type KelompokPendapatanBelanja struct {
	Kode       string `validate:"required,min=1,max=10" json:"kode"`
	Keterangan string `validate:"required" json:"keterangan"`
}

func (req *KelompokPendapatanBelanja) ToModelKelompokPendapatan() *models.KelompokPendapatan {
	return &models.KelompokPendapatan{
		Kode:       req.Kode,
		Keterangan: req.Keterangan,
	}
}
func (req *KelompokPendapatanBelanja) ToModelKelompokPendapatanCsv() models.KelompokPendapatan {
	return models.KelompokPendapatan{
		Kode:       req.Kode,
		Keterangan: req.Keterangan,
	}
}
