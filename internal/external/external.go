package external

//go:generate mockgen -source=external.go -destination=../mock/external.go -package=mock

import (
	"log/slog"
	"time"

	"github.com/go-resty/resty/v2"
)

type IExternal interface {
	ListUsers() []ResponseUsers
}

type External struct {
	getUser *resty.Client
}

func NewExternal() *External {
	getUser := resty.
		New().
		SetBaseURL("https://gorest.co.in").
		SetTimeout(15 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(10 * time.Second).
		SetRetryMaxWaitTime(60 * time.Second)

	getUser.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		slog.InfoContext(
			request.Context(),
			"request GET /public/v2/users",
			slog.String("method", request.Method),
			slog.String("url", request.URL),
			slog.Any("headers", request.Header),
			slog.Any("body", request.Body),
		)
		return nil
	})

	getUser.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		slog.InfoContext(
			response.Request.Context(),
			"response GET /public/v2/users",
			slog.Int("status_code", response.StatusCode()),
			slog.String("body", response.String()),
		)
		return nil
	})

	return &External{
		getUser: getUser,
	}
}

func (e *External) ListUsers() []ResponseUsers {
	var payload []ResponseUsers
	res, err := e.
		getUser.
		R().
		SetResult(&payload).
		Get("/public/v2/users")
	if err != nil {
		panic(err)
	}
	if res.IsError() {
		panic(res.Error())
	}

	return payload
}

type ResponseUsers struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}
