// package ini berisi berbagai macam konfigurasi proyek seperti koneksi ke database dan migrasi database/table
package configs

import (
	"fmt"
	"os"
	"uji_keterampilan/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// variabel global yang dipanggil di package lain
var DB *gorm.DB

// struct konfigurasi database yang nantinya diisi dari file .env
type DBconfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// func untuk inisialisasi konfigurasi koneksi ke database mysql
func InitDB() {
	var dbConfig = DBconfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database error")
	}
	migration()
}

// func untuk auto migrasi jika belum ada table yang dibuat
func migration() {
	DB.AutoMigrate(&models.Films{})
}
