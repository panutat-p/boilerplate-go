package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"boilerplate-go/internal/store"
	"boilerplate-go/internal/usecase"
)

type Handler struct {
	useCase usecase.IUseCase
	store   store.IStore
}

func NewHandler(useCase usecase.IUseCase, store store.IStore) *Handler {
	return &Handler{
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
	fruits, err := h.useCase.GetFruits(nil)
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
