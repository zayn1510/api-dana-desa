package requests

import "apidanadesa/app/models"

type KegiatanRequestCreate struct {
	IdBidang     uint   `json:"id_bidang" validate:"required"`
	IdSubBidang  uint   `json:"id_sub_bidang" validate:"required"`
	Keterangan   string `json:"keterangan" validate:"required"`
	KodeKegiatan string `json:"kode_kegiatan" validate:"required"`
}

func (req *KegiatanRequestCreate) ToModelkegiatan() *models.Kegiatan {
	return &models.Kegiatan{
		IdBidang:     req.IdBidang,
		IdSubBidang:  req.IdSubBidang,
		Keterangan:   req.Keterangan,
		KodeKegiatan: req.KodeKegiatan,
	}
}
