package v1

import (
	"net/http"
	"strings"
)

type Error struct {
	Code   int                    `json:"code"`
	Errors map[string]interface{} `json:"erorrs"`
}

// func newError(err error) Error {
// 	e := Error{}
// 	e.Errors = make(map[string]interface{})
// 	switch v := err.(type) {
// 	case *echo.HTTPError:
// 		e.Code = http.StatusInternalServerError
// 		e.Errors["message"] = v.Message
// 		e.Errors["data"] = nil
// 	default:
// 		e.Code = http.StatusInternalServerError
// 		e.Errors["message"] = v.Error()
// 		e.Errors["type"] = v
// 		e.Errors["data"] = nil
// 	}

// 	return e
// }

func notFound() Error {
	e := Error{}
	e.Code = http.StatusNotFound
	e.Errors = make(map[string]interface{})
	e.Errors["message"] = "resource not found"

	return e
}

func badRequest() Error {
	e := Error{}
	e.Code = http.StatusBadRequest
	e.Errors = make(map[string]interface{})
	e.Errors["message"] = "bad request"

	return e
}

func entityError(err error) Error {
	e := Error{}
	e.Code = http.StatusUnprocessableEntity
	e.Errors = make(map[string]interface{})
	e.Errors["message"] = strings.Split(strings.Split(err.Error(), ", ")[1], "=")[1]

	return e
}
