package main

import (
	"log/slog"
	"os"

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
	logger := slog.New(slog.NewTextHandler(os.Stdin, nil))
	slog.SetDefault(logger)

	var conf config.Config
	err := env.Parse(&conf)
	if err != nil {
		panic(err)
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate = pkg.RegisterNullTypes(validate)
	err = validate.Struct(conf)
	if err != nil {
		panic(err)
	}

	pkg.PrintJSON(conf)

	st := store.NewStore(&conf)
	uc := usecase.NewUseCase(&conf, st)
	h := handler.NewHandler(&conf, uc, st)

	e := echo.New()

	e.GET("/", h.Health)
	e.GET("/public/fruits", h.GetFruits)

	err = e.Start(":" + conf.Port)
	if err != nil {
		panic(err)
	}
}
