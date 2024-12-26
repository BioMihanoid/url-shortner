package handler

import (
	"github.com/BioMihanoid/url-shortner/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	url := router.Group("/url")
	{
		url.POST("/short", h.SaveUrl)
		url.GET("/resolve", h.GetUrl)
	}

	return router
}
