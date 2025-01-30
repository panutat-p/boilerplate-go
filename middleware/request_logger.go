package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			r := c.Request()

			if r.URL.Path == "/" || r.URL.Path == "/healthz" || r.URL.Path == "/metrics" || r.URL.Path == "/readyz" {
				return next(c)
			}

			if r.Body != nil {
				b, _ := io.ReadAll(r.Body)
				var removedSpacesBody []byte
				for _, c := range b {
					if c != ' ' && c != '\n' && c != '\r' && c != '\t' {
						removedSpacesBody = append(removedSpacesBody, c)
					}
				}
				fmt.Printf("%s    %s %s %s\n", time.Now().Format(time.DateTime), r.Method, r.URL.Path, string(removedSpacesBody))
				r.Body = io.NopCloser(bytes.NewBuffer(removedSpacesBody))
			}

			return next(c)
		}
	}
}
