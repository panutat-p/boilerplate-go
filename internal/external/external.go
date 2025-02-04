package external

//go:generate mockgen -source=external.go -destination=../mock/external.go -package=mock

import (
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
		SetTimeout(10 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)

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
