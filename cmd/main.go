package main

import (
	"log/slog"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"boilerplate-go/config"
	"boilerplate-go/handler"
	"boilerplate-go/internal/external"
	"boilerplate-go/internal/store"
	"boilerplate-go/internal/usecase"
	"boilerplate-go/middleware"
	"boilerplate-go/pkg"
)

func main() {
	logger := slog.New(
		slog.NewTextHandler(
			os.Stdin,
			&slog.HandlerOptions{
				Level:     slog.LevelInfo,
				AddSource: true,
			},
		),
	)
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
	ex := external.NewExternal()
	uc := usecase.NewUseCase(&conf, st, ex)
	h := handler.NewHandler(&conf, uc, st)

	e := echo.New()
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORSWithConfig(
		echomiddleware.CORSConfig{
			Skipper:      echomiddleware.DefaultSkipper,
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{}, // allow all
		}),
	)
	e.Use(middleware.RequestLogger())

	e.GET("/", h.Health)
	e.GET("/public/fruits", h.GetFruits)

	err = e.Start(":" + conf.Port)
	if err != nil {
		panic(err)
	}
}
