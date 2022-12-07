// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware

	// Swagger docs.
	"github.com/rmscoal/Authenticator-API/internal/usecase"
	"github.com/rmscoal/Authenticator-API/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Authenticator API made with Golang
// @description An Authenticator API to generate token
// @version     1.0
// @host        localhost:8081
// @BasePath    /v1
func NewRouter(handler *echo.Echo, l logger.Interface, t usecase.User) {
	// Options
	log := logrus.New()
	handler.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			switch {
			case values.Status >= 200 && values.Status < 400:
				log.WithFields(logrus.Fields{
					"Latency": values.Latency,
					"URI":     values.URI,
					"status":  values.Status,
				}).Info("Success")
			case values.Status >= 400 && values.Status < 500:
				log.WithFields(logrus.Fields{
					"Latency": values.Latency,
					"URI":     values.URI,
					"status":  values.Status,
				}).Warn("Bad Request")
			default:
				log.WithFields(logrus.Fields{
					"Latency": values.Latency,
					"URI":     values.URI,
					"status":  values.Status,
				}).Log(2, "Internal Error")
			}
			return nil
		},
	}))
	handler.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	handler.Validator = &CustomValidator{Validator: validator.New()}

	// Swagger
	handler.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routers
	h := handler.Group("/v1")
	{
		newAuthenticatorRoutes(h, t, l)
	}
}
