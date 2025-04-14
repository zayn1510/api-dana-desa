package requests

import "mime/multipart"

type PerangkatDesaCreateRequest struct {
	IDJabatan    uint                  `form:"id_jabatan" validate:"required"`
	NamaLengkap  string                `form:"nama_lengkap" validate:"required,max=100"`
	TempatLahir  string                `form:"tempat_lahir" validate:"required,max=100"`
	TglLahir     string                `form:"tgl_lahir" validate:"required"`
	JenisKelamin string                `form:"jenis_kelamin" validate:"required,oneof='Laki-laki' 'Perempuan'"`
	NoSK         string                `form:"no_sk" validate:"required,max=100"`
	TglSK        string                `form:"tgl_sk" validate:"required"` // parse manual juga
	NoHandphone  string                `form:"no_handphone" validate:"required,max=20"`
	Foto         *multipart.FileHeader `form:"foto" validate:"omitempty"`
}
