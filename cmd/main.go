package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

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

	pkg.PrintJSON(conf)

	validate := validator.New(validator.WithRequiredStructEnabled())
	validate = pkg.RegisterNullTypes(validate)
	validate = pkg.RegisterDecimalTypes(validate)
	err = validate.Struct(conf)
	if err != nil {
		panic(err)
	}

	st := store.NewStore(&conf)
	uc := usecase.NewUseCase(&conf, st)
	h := handler.NewHandler(&conf, uc, st)

	e := echo.New()
	e.Use(echomiddleware.CORSWithConfig(
		echomiddleware.CORSConfig{
			Skipper:      echomiddleware.DefaultSkipper,
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		}),
	)

	e.GET("/", h.Health)
	e.GET("/public/fruits", h.GetFruits)

	err = e.Start(":" + conf.Port)
	if err != nil {
		panic(err)
	}
}
