package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rmscoal/Authenticator-API/internal/entity"
	"github.com/rmscoal/Authenticator-API/internal/usecase"
	"github.com/rmscoal/Authenticator-API/pkg/logger"
	"github.com/rmscoal/Authenticator-API/pkg/tokenizer"
)

type authenticatorRoutes struct {
	u usecase.User
	l logger.Interface
	t tokenizer.Tokenizer
}

func newAuthenticatorRoutes(handler *echo.Group, u usecase.User, l logger.Interface, t tokenizer.Tokenizer) {
	r := &authenticatorRoutes{u, l, t}

	h := handler.Group("/credential")
	{
		h.POST("/login", r.login)
	}
}

type loginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type loginResponse struct {
	Status string   `json:"status"`
	User   userData `json:"user"`
	Token  []byte   `json:"token"`
}

type userData struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (r *authenticatorRoutes) login(c echo.Context) error {
	body := &loginRequest{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, badRequest())
	}

	r.l.Info("http - v1 - login - validating")
	if err := c.Validate(body); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, entityError(err))
	}

	r.l.Info("http - v1 - login - querying user")
	user, err := r.u.Find(c.Request().Context(),
		entity.User{
			Username: body.Username,
			Password: body.Password,
		},
	)
	// If an error is generated from r.t.Find() it is caused
	// by query result producing no row.
	if err != nil {
		r.l.Error(err, "http - v1 - login")
		return c.JSON(http.StatusNotFound, notFound())
	}

	token, err := r.t.GenerateFromPassword([]byte(body.Password))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newError(err))
	}

	return c.JSON(http.StatusOK, loginResponse{
		Status: "success",
		User: userData{
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Token: token,
	})
}
