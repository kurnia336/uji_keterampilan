package controllers

import (
	"net/http"
	"strconv"
	config "uji_keterampilan/configs"
	"uji_keterampilan/models"

	"github.com/labstack/echo/v4"
)

// berfungsi untuk meng fetch all/select all data dari table films di database
func FilmsController(c echo.Context) error {

	var films []models.Films
	//konek db dan query select all table films
	result := config.DB.Find(&films)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    films,
	})
}

// meng-fetch satu data dari table
func DetailFilmsController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	var films models.Films
	//konek db dan query data dengan kondisi where id
	result := config.DB.Find(&films, "id = ?", id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    films,
	})
}

// fungsi meng-insert satu data pada table films
func AddFilmsController(c echo.Context) error {
	var films models.Films
	c.Bind(&films)

	//konek db dan query insert
	result := config.DB.Create(&films)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    films,
	})
}

// fungsi untuk update satu data dari table dengan kondisi where id
func UpdateFilmsController(c echo.Context) error {
	var films models.Films
	var id, _ = strconv.Atoi(c.Param("id"))

	title := c.FormValue("title")
	year := c.FormValue("year")
	years, _ := strconv.Atoi(year)
	rating := c.FormValue("rating")
	ratings, _ := strconv.Atoi(rating)
	produser := c.FormValue("produser")
	synopsis := c.FormValue("synopsis")

	// konek db dan query data dari table dengan kondisi where id
	result := config.DB.Find(&films, "id = ?", id)
	// update data dari form input
	films.Title = title
	films.Year = years
	films.Rating = ratings
	films.Produser = produser
	films.Synopsis = synopsis
	// simpan data ke db
	config.DB.Save(&films)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    films,
	})
}

// fungsi untuk menghapus satu data dari table dengan kondisi where id
func DeleteFilmsController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var films models.Films
	// konek db dan query ambil satu data dari table dengan kondisi where id
	result := config.DB.Find(&films, "id = ?", id)
	// query delete satu data
	config.DB.Delete(&films)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Error",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success",
		Data:    films,
	})
}
