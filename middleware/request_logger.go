package middleware

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			r := c.Request()

			if r.URL.Path == "/" || r.URL.Path == "/health" {
				return next(c)
			}

			if r.Body != nil {
				b, err := io.ReadAll(r.Body)
				if err != nil {
					r.Body = io.NopCloser(bytes.NewBuffer(nil))
					return next(c)
				}
				r.Body = io.NopCloser(bytes.NewBuffer(b))

				text := strings.ReplaceAll(string(b), "\n", "")

				fmt.Printf("%s    %s %s %s\n", time.Now().Format(time.DateTime), r.Method, r.URL.Path, text)
			}

			return next(c)
		}
	}
}
