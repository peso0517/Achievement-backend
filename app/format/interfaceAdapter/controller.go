package controller

import (
	model "Achievement/backend/app/format/domain"
	Ucase "Achievement/backend/app/format/usecase"
	"net/http"
	"log"

	"github.com/labstack/echo/v4"
)

// Create is..
func Create(c echo.Context) error {
	formatmodel := new(model.ReceiveFormatData)
	if err := c.Bind(formatmodel); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	log.Println(formatmodel, "interface--------data----------")
	// type usecase.FomatUsecase is usecase method
	formatusecase := Ucase.NewFormatUsecase()
	err1 := formatusecase.Create(formatmodel)
	if err1 != nil {
		log.Println(err1, "create---------")
		return c.NoContent(http.StatusBadRequest)
	}
	return c.NoContent(http.StatusOK)
}
