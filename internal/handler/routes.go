package handler

import "github.com/gin-gonic/gin"

func (h *Handler) WeatherRoutes(r *gin.Engine) {
	r.GET("/weather/:location", h.getWeather)
	r.PUT("/weather/:location", h.updateWeather)
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()
	h.WeatherRoutes(router)
	return router
}
