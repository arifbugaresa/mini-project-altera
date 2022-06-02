package api

import echo "github.com/labstack/echo/v4"

// APIController register all v1 API with routing path
func APIController(e *echo.Echo)  {

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

}