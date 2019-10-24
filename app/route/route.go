package route

import (
	format "Achievement/backend/app/format/interfaceAdapter"

	"github.com/labstack/echo/v4"
)

// Route is...
func Route(e *echo.Echo) {
	e.POST("/format", format.Create)
}
