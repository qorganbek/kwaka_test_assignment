package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kwaka_test/internal/entity"
	"net/http"
)

func (h *Handler) getWeather(ctx *gin.Context) {
	location := ctx.Param("location")
	var weather entity.Weather
	weather, err := h.service.GetWeather(location)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSONP(http.StatusOK, weather)
}

func (h *Handler) updateWeather(ctx *gin.Context) {
	location := ctx.Param("location")

	weather, err := h.service.GetWeather(location)
	fmt.Println("Weather from update ", weather)
	if err != nil {
		weather, err = h.service.CreateWeather(location)
		ctx.JSON(http.StatusOK, weather)
		return
	}

	err = h.service.UpdateWeather(location, entity.UpdateWeather{
		Temperature: weather.Temperature,
		Description: weather.Description,
		FeelsLike:   weather.FeelsLike,
	})

	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, weather)
}
