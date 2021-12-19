package routes

import (
	"server/web/api"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	g := e.Group("/api")
	{
		g.GET("/popular", api.FetchMostPopularVideos())
	}
}
