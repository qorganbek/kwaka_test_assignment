package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/qorganbek/kwaka_test_assignment/internal/entity"
	"github.com/qorganbek/kwaka_test_assignment/internal/service"
	mock_service "github.com/qorganbek/kwaka_test_assignment/internal/service/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_getWeather(t *testing.T) {
	type mockBehavior func(s *mock_service.MockWeather, weather entity.Weather)

	testTable := []struct {
		name                 string
		inputBody            string
		inputWeather         entity.Weather
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: "",
			inputWeather: entity.Weather{
				ID:          10,
				Location:    "Seattle",
				Temperature: 6.3,
				Description: "light rain",
				FeelsLike:   4.82,
			},
			mockBehavior: func(s *mock_service.MockWeather, weather entity.Weather) {
				s.EXPECT().GetWeather(weather.Location).Return(entity.Weather{
					ID:          11,
					Location:    "London",
					Temperature: 12.1,
					Description: "scattered clouds",
					FeelsLike:   11.43,
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":10,\"location\":\"Seattle\",\"temp\":6.3,\"description\":\"light rain\",\"feels_like\":4.82}",
		},
		{
			name:      "OK",
			inputBody: "",
			inputWeather: entity.Weather{
				ID:          12,
				Location:    "Talgar",
				Temperature: -17.91,
				Description: "light snow",
				FeelsLike:   -23.32,
			},
			mockBehavior: func(s *mock_service.MockWeather, weather entity.Weather) {
				s.EXPECT().GetWeather(weather.Location).Return(entity.Weather{
					ID:          10,
					Location:    "Seattle",
					Temperature: 6.3,
					Description: "light rain",
					FeelsLike:   4.82,
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":10,\"location\":\"Seattle\",\"temp\":6.3,\"description\":\"light rain\",\"feels_like\":4.82}",
		},
		{
			name:      "OK",
			inputBody: "",
			inputWeather: entity.Weather{
				ID:          10,
				Location:    "Seattle",
				Temperature: 6.3,
				Description: "light rain",
				FeelsLike:   4.82,
			},
			mockBehavior: func(s *mock_service.MockWeather, weather entity.Weather) {
				s.EXPECT().GetWeather(weather.Location).Return(entity.Weather{
					ID:          10,
					Location:    "Seattle",
					Temperature: 6.3,
					Description: "light rain",
					FeelsLike:   4.82,
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"id\":10,\"location\":\"Seattle\",\"temp\":6.3,\"description\":\"light rain\",\"feels_like\":4.82}",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			weather := mock_service.NewMockWeather(c)
			testCase.mockBehavior(weather, testCase.inputWeather)

			services := &service.Service{Weather: weather}
			handler := New(services)

			r := gin.New()
			r.GET("/weather/"+testCase.inputWeather.Location, handler.getWeather)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/weather/"+testCase.inputWeather.Location,
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
