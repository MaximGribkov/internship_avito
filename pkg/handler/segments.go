package handler

import (
	"github.com/gin-gonic/gin"
	"internship_avito/pkg/model"
	"net/http"
)

func (h *Handler) createSegments(c *gin.Context) {
	var input model.Segments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	id, err := h.services.LogicSegment.CreateSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteSegments(c *gin.Context) {
	var input model.Segments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	answer, err := h.services.LogicSegment.DeleteSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"answer": answer,
	})
}

// Получение списка активных сегментов пользователя
func (h *Handler) getSegmentsUser(c *gin.Context) {

}
