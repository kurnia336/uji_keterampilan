// package routes berfungsi untuk mengatur pengalamatan dari setiap endpoint
package routes

import (
	"uji_keterampilan/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.GET("/films", controllers.FilmsController)              //endpoint READ method
	e.GET("/films/:id", controllers.DetailFilmsController)    //endpoint DETAIL method
	e.POST("/films", controllers.AddFilmsController)          //endpoint CREATE method
	e.PUT("/films/:id", controllers.UpdateFilmsController)    //endpoint UPDATE method
	e.DELETE("/films/:id", controllers.DeleteFilmsController) //endpoint DELETE method
	return e
}
