package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}

//validasi struct dengan menggunakan library go-playground/validator/v10.
// Dalam hal ini, kita akan memanggil fungsi ValidateStruct() dengan parameter
//sebuah struct yang ingin divalidasi. Fungsi ini akan mengembalikan
//error jika terdapat kesalahan dalam validasi, atau nil jika tidak ada kesalahan.
