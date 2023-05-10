package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}

//Kode tersebut berfungsi untuk membuat koneksi ke database MySQL menggunakan driver gorm.io/driver/mysql.
//Dalam hal ini, kita akan memanggil fungsi NewDBConnection() dengan parameter DSN (Data Source Name)
//yang berisi informasi host, port, nama database, username, dan password. Fungsi ini akan
//mengembalikan instance *gorm.DB yang bisa digunakan untuk menjalankan query atau migrasi database.
