package main

import (
	"server/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

//var e = createMux()

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	// e.GET("/", articleIndex)

	// e.Logger.Fatal(e.Start(":8080"))
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.Init(e)
	e.Logger.Fatal(e.Start(":8080"))

}

// func createMux() *echo.Echo {
// 	e := echo.New()

// 	e.Use(middleware.Recover())
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Gzip())

// 	return e
// }

// func articleIndex(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
