package handler

import (
	"log/slog"
	"net/http"

	govalidator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"boilerplate-go/config"
	"boilerplate-go/internal/store"
	"boilerplate-go/internal/usecase"
	"boilerplate-go/pkg"
)

type Handler struct {
	config    *config.Config
	useCase   usecase.IUseCase
	store     store.IStore
	validator *govalidator.Validate
}

func NewHandler(config *config.Config, useCase usecase.IUseCase, store store.IStore) *Handler {

	validator := govalidator.New(govalidator.WithRequiredStructEnabled())
	validator = pkg.RegisterNullTypes(validator)

	return &Handler{
		config:    config,
		useCase:   useCase,
		store:     store,
		validator: validator,
	}
}

func (h *Handler) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"message": "API is running",
		"version": "1.0.0",
	})
}

func (h *Handler) GetFruits(c echo.Context) error {
	ctx := c.Request().Context()
	slog.Info("GET /public/fruits")

	fruits, err := h.useCase.GetFruits(ctx)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]any{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		fruits,
	)
}
