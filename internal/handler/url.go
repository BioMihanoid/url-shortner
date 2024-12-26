package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SaveUrl(c *gin.Context) {
	var link string

	if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sUrl, err := h.service.URL.SaveUrl(link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": sUrl,
	})
}

func (h *Handler) GetUrl(c *gin.Context) {
	var link string

	if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lUrl, err := h.service.URL.GetLongUrl(link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url": lUrl,
	})
}
