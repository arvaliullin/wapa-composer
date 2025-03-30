package handlers

import (
	"fmt"
	"net/http"

	"github.com/arvaliullin/wapa-composer/internal/delivery"
	"github.com/arvaliullin/wapa-composer/internal/domain"
	"github.com/arvaliullin/wapa-composer/internal/persistence"
	"github.com/labstack/echo/v4"
)

type DesignHandler struct {
	repo persistence.DesignRepositoryContract
}

func RegisterDesignHandler(httpService delivery.HttpService, repo persistence.DesignRepositoryContract) {
	handler := &DesignHandler{repo: repo}
	e := httpService.(*delivery.EchoHttpService).Echo
	e.GET("/api/designs", handler.GetAll)
	e.POST("/api/design", handler.Create)
	e.DELETE("/api/design/:id", handler.Delete)
}

// GetAll получает список планов экспериментов
// @Summary получает список планов экспериментов
// @Description Возвращает список планов экспериментов.
// @Tags Design
// @Produce json
// @Success 200 {array} domain.Design
// @Failure 500 {object} string "Ошибка сервера"
// @Router /api/designs [get]
func (h *DesignHandler) GetAll(c echo.Context) error {
	designs, err := h.repo.GetAll()

	if err != nil {
		msg := fmt.Sprintf("Ошибка получения планов экспериментов: %v", err)
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, designs)
}

// Create создает новый план эксперимента
// @Summary создает новый план эксперимента
// @Description Создает новый план эксперимента
// @Tags Design
// @Accept json
// @Produce json
// @Param design body domain.Design true "Данные плана эксперимента"
// @Success 201 {object} string "Создано"
// @Failure 400 {object} string "Неверный запрос"
// @Failure 500 {object} string "Ошибка сервера"
// @Router /api/design [post]
func (h *DesignHandler) Create(c echo.Context) error {
	var design domain.Design
	if err := c.Bind(&design); err != nil {
		msg := fmt.Sprintf("Неверный запрос: %v", err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	if err := h.repo.Create(design); err != nil {
		msg := fmt.Sprintf("Ошибка сервера: %v", err)
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.JSON(http.StatusCreated, "Создано")
}

// Delete удаляет план эксперимента
// @Summary удаляет план эксперимента
// @Description Удаляет план эксперимента по ID.
// @Tags Design
// @Produce json
// @Param id path string true "ID плана эксперимента"
// @Success 200 {object} string "Удалено"
// @Failure 400 {object} string "Неверный запрос"
// @Failure 500 {object} string "Ошибка сервера"
// @Router /api/design/{id} [delete]
func (h *DesignHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "Неверный запрос")
	}
	if err := h.repo.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, "Ошибка сервера")
	}
	return c.JSON(http.StatusOK, "Удалено")
}
