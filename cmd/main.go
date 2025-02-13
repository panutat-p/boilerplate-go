package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

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

var (
	stop = make(chan os.Signal, 1)
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
	e.GET("/public/reflect", h.GetReflect)
	e.GET("/public/fruits", h.GetFruits)

	go func() {
		err := e.Start(":" + conf.Port)
		if err != nil {
			slog.Error("Echo server is closed", "reason", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		slog.ErrorContext(
			ctx,
			"Server forced to shutdown:",
			slog.Any("error", err),
		)
	}

	slog.Info("Server shutdown successfully")
}
