package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"boilerplate-go/internal/config"
	"boilerplate-go/internal/store"
	"boilerplate-go/internal/usecase"
)

type Handler struct {
	config  *config.Config
	useCase usecase.IUseCase
	store   store.IStore
}

func NewHandler(config *config.Config, useCase usecase.IUseCase, store store.IStore) *Handler {
	return &Handler{
		config:  config,
		useCase: useCase,
		store:   store,
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
