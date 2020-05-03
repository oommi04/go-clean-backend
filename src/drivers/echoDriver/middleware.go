package echoDriver

import "github.com/labstack/echo"

type GoMiddleware struct {

}

func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}