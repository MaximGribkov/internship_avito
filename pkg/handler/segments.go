package handler

import (
	"github.com/gin-gonic/gin"
	"internship_avito/pkg/model"
	"net/http"
)

// Метод создания сегмента
func (h *Handler) createSegments(c *gin.Context) {
	var input model.Segments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	answer, err := h.services.LogicSegment.CreateSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"segment_name": answer,
	})
}

// Метод удаления сегмента
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

// Метод получения количества пользователей для одного сегмента
func (h *Handler) userCountInSegment(c *gin.Context) {
	var input model.Segments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	answer, err := h.services.LogicSegment.UserCountInSegment(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"count": answer,
	})
}
