package internal

import (
	"github.com/labstack/echo"
	"github.com/yuuLab/go-rest-api/internal/presentation"
)

func RunServer() {
	e := echo.New()
	presentation.Router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
