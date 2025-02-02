package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
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
	validator *validator.Validate
}

func NewHandler(config *config.Config, useCase usecase.IUseCase, store store.IStore) *Handler {

	val := validator.New(validator.WithRequiredStructEnabled())
	val = pkg.RegisterNullTypes(val)
	val = pkg.RegisterDecimalTypes(val)

	return &Handler{
		config:    config,
		useCase:   useCase,
		store:     store,
		validator: val,
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
