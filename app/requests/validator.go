package requests

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Global validator instance
var validate = validator.New()

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Validate(request interface{}) ([]ValidationErrorResponse, error) {
	err := validate.Struct(request)
	if err == nil {
		return nil, nil
	}

	var validationErrors []ValidationErrorResponse
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			// Buat pesan error lebih deskriptif
			message := getErrorMessage(e)
			validationErrors = append(validationErrors, ValidationErrorResponse{
				Field:   strings.ToLower(e.Field()),
				Message: strings.ToLower(message),
			})
		}
	}

	return validationErrors, err
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("Field %s wajib diisi", e.Field())
	case "email":
		return fmt.Sprintf("Field %s harus berupa email yang valid", e.Field())
	case "numeric":
		return fmt.Sprintf("Field %s harus berupa angka", e.Field())
	case "min":
		return fmt.Sprintf("Field %s harus memiliki minimal %s karakter", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("Field %s tidak boleh lebih dari %s karakter", e.Field(), e.Param())
	case "len":
		return fmt.Sprintf("Field %s harus memiliki tepat %s karakter", e.Field(), e.Param())
	case "eq":
		return fmt.Sprintf("Field %s harus sama dengan %s", e.Field(), e.Param())
	case "ne":
		return fmt.Sprintf("Field %s tidak boleh sama dengan %s", e.Field(), e.Param())
	case "lt":
		return fmt.Sprintf("Field %s harus kurang dari %s", e.Field(), e.Param())
	case "lte":
		return fmt.Sprintf("Field %s harus kurang dari atau sama dengan %s", e.Field(), e.Param())
	case "gt":
		return fmt.Sprintf("Field %s harus lebih dari %s", e.Field(), e.Param())
	case "gte":
		return fmt.Sprintf("Field %s harus lebih dari atau sama dengan %s", e.Field(), e.Param())
	case "alpha":
		return fmt.Sprintf("Field %s hanya boleh berisi huruf", e.Field())
	case "alphanum":
		return fmt.Sprintf("Field %s hanya boleh berisi huruf dan angka", e.Field())
	case "boolean":
		return fmt.Sprintf("Field %s harus berupa true atau false", e.Field())
	case "url":
		return fmt.Sprintf("Field %s harus berupa URL yang valid", e.Field())
	case "uuid":
		return fmt.Sprintf("Field %s harus berupa UUID yang valid", e.Field())
	case "uuid4":
		return fmt.Sprintf("Field %s harus berupa UUID versi 4 yang valid", e.Field())
	case "uuid5":
		return fmt.Sprintf("Field %s harus berupa UUID versi 5 yang valid", e.Field())
	case "ipv4":
		return fmt.Sprintf("Field %s harus berupa alamat IPv4 yang valid", e.Field())
	case "ipv6":
		return fmt.Sprintf("Field %s harus berupa alamat IPv6 yang valid", e.Field())
	case "ip":
		return fmt.Sprintf("Field %s harus berupa alamat IP yang valid", e.Field())
	case "contains":
		return fmt.Sprintf("Field %s harus mengandung %s", e.Field(), e.Param())
	case "excludes":
		return fmt.Sprintf("Field %s tidak boleh mengandung %s", e.Field(), e.Param())
	case "containsany":
		return fmt.Sprintf("Field %s harus mengandung setidaknya satu karakter dari %s", e.Field(), e.Param())
	case "excludesall":
		return fmt.Sprintf("Field %s tidak boleh mengandung karakter dari %s", e.Field(), e.Param())
	case "oneof":
		return fmt.Sprintf("Field %s harus salah satu dari %s", e.Field(), e.Param())
	case "startswith":
		return fmt.Sprintf("Field %s harus diawali dengan %s", e.Field(), e.Param())
	case "endswith":
		return fmt.Sprintf("Field %s harus diakhiri dengan %s", e.Field(), e.Param())
	case "ascii":
		return fmt.Sprintf("Field %s hanya boleh berisi karakter ASCII", e.Field())
	case "printascii":
		return fmt.Sprintf("Field %s hanya boleh berisi karakter ASCII yang dapat dicetak", e.Field())
	case "multibyte":
		return fmt.Sprintf("Field %s harus berisi karakter multibyte", e.Field())
	case "datauri":
		return fmt.Sprintf("Field %s harus berupa Data URI yang valid", e.Field())
	case "base64":
		return fmt.Sprintf("Field %s harus berupa string base64 yang valid", e.Field())
	case "hexadecimal":
		return fmt.Sprintf("Field %s harus berupa string heksadesimal yang valid", e.Field())
	case "hexcolor":
		return fmt.Sprintf("Field %s harus berupa kode warna heksadesimal yang valid", e.Field())
	case "lowercase":
		return fmt.Sprintf("Field %s hanya boleh berisi huruf kecil", e.Field())
	case "uppercase":
		return fmt.Sprintf("Field %s hanya boleh berisi huruf besar", e.Field())
	case "datetime":
		return fmt.Sprintf("Field %s harus sesuai format waktu %s", e.Field(), e.Param())
	case "json":
		return fmt.Sprintf("Field %s harus berupa JSON yang valid", e.Field())
	case "uuid3":
		return fmt.Sprintf("Field %s harus berupa UUID versi 3 yang valid", e.Field())
	case "md5":
		return fmt.Sprintf("Field %s harus berupa hash MD5 yang valid", e.Field())
	case "sha256":
		return fmt.Sprintf("Field %s harus berupa hash SHA256 yang valid", e.Field())
	case "sha512":
		return fmt.Sprintf("Field %s harus berupa hash SHA512 yang valid", e.Field())
	case "mac":
		return fmt.Sprintf("Field %s harus berupa alamat MAC yang valid", e.Field())
	case "fqdn":
		return fmt.Sprintf("Field %s harus berupa Fully Qualified Domain Name (FQDN)", e.Field())
	case "unique":
		return fmt.Sprintf("Field %s harus unik", e.Field())
	default:
		return fmt.Sprintf("Field %s tidak valid", e.Field())
	}
}
