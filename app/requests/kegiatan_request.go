package requests

import "apidanadesa/app/models"

type KegiatanRequestCreate struct {
	IdBidang     uint   `json:"id_bidang" binding:"required"`
	IdSubBidang  uint   `json:"id_sub_bidang" binding:"required"`
	Keterangan   string `json:"keterangan" binding:"required"`
	KodeKegiatan string `json:"kode_kegiatan" binding:"required"`
}

func (req *KegiatanRequestCreate) ToModelkegiatan() *models.Kegiatan {
	return &models.Kegiatan{
		IdBidang:     req.IdBidang,
		IdSubBidang:  req.IdSubBidang,
		Keterangan:   req.Keterangan,
		KodeKegiatan: req.KodeKegiatan,
	}
}
