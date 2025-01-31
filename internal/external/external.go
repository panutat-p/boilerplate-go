package external

//go:generate mockgen -source=external.go -destination=mock_external/external.go -package=mock_external

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type IExternal interface{}

type External struct {
	resty *resty.Client
}

func NewExternal() *External {
	r := resty.
		New().
		SetTimeout(10 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second)

	return &External{
		resty: r,
	}
}
