package main

import (
	"log"
	"uji_keterampilan/configs"
	"uji_keterampilan/routes"

	//library pengaturan env
	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv() //load environmentnya
	configs.InitDB() //konfigurasi dan koneksi db
	e := echo.New() //inisialisasi pertama framework echo
	e = routes.InitRoute(e) //inisialisasi route
	e.Start(":8000") //start di port
}

//func meng load env menggunakan library
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
