package models

type Films struct {
	Id      	int    	`json:"id"`
	Title   	string 	`json:"title"`
	Year		int		`json:"year"`
	Rating 		int		`json:"rating"`
	Produser 	string  `json:"produser"`
	Synopsis 	string  `json:"synopsis"`
}