package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"boilerplate-go/handler"
	"boilerplate-go/internal/config"
	"boilerplate-go/internal/store"
	"boilerplate-go/internal/usecase"
	"boilerplate-go/pkg"
)

func main() {
	var conf config.Config
	err := env.Parse(&conf)
	if err != nil {
		panic(err)
	}

	validate := validator.New()
	err = validate.Struct(conf)
	if err != nil {
		panic(err)
	}

	pkg.PrintJSON(conf)

	st := store.NewStore()
	uc := usecase.NewUseCase(st)
	h := handler.NewHandler(uc, st)

	e := echo.New()

	e.GET("/", h.Health)
	e.GET("/public/fruits", h.GetFruits)

	err = e.Start(":" + conf.Port)
	if err != nil {
		panic(err)
	}
}
