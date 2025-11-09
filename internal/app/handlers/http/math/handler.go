package math

import (
	"fmt"
	"github.com/TemaKut/tt-perx/internal/app/handlers/http/math/structs"
	mathdto "github.com/TemaKut/tt-perx/internal/dto/math"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandleArithmeticProgressionTasksAdd(c echo.Context) error {
	var params structs.ArithmeticProgressionTaskAdd

	if err := c.Bind(&params); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("error bind params. %s", err))
	}

	h.service.AddArithmeticProgressionTask(decodeArithmeticProgressionTaskAdd(params))

	return nil
}

func (h *Handler) HandleArithmeticProgressionTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, encodeArithmeticProgressionTasks(h.service.ArithmeticProgressionTasks()))
}

func encodeArithmeticProgressionTasks(tasks []mathdto.ArithmeticProgressionTask) []structs.ArithmeticProgressionTask {
	result := make([]structs.ArithmeticProgressionTask, 0, len(tasks))

	for _, task := range tasks {
		result = append(result, encodeArithmeticProgressionTask(task))
	}

	return result
}
