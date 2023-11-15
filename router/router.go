package router

import (
	"KobaCareer_API/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter(ic controller.IInternshipController) *echo.Echo {
	e := echo.New()
	e.GET("/internships", ic.GetAllInternships)
	e.GET("/internships/:internshipId", ic.GetInternshipById)
	e.POST("/internships", ic.CreateInternship)
	e.PUT("/internships/:internshipId", ic.UpdateInternship)
	e.DELETE("/internships/:internshipId", ic.DeleteInternship)
	return e
}
