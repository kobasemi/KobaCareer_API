package main

import (
	"KobaCareer_API/controller"
	"KobaCareer_API/db"
	_ "KobaCareer_API/migrate"
	"KobaCareer_API/repository"
	"KobaCareer_API/router"
	"KobaCareer_API/usecase"
	"KobaCareer_API/validator"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := db.NewDB()
	internValidate := validator.NewInternshipValidator()
	internRepository := repository.NewInternshipRepository(db)
	internUsecase := usecase.NewInternshipUsecase(internRepository, internValidate)
	internController := controller.NewInternshipController(internUsecase)
	e := router.NewRouter(internController)
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(":8080"))
}
