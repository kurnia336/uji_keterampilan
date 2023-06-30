//package model yang berisi struct untuk mengatur respon yang dihasilkan dari setiap operasi yang dijalankan
package models

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"` //interface empty
}