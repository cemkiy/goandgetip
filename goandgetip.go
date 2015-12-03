package main

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/jokerkart/paladin/utils"
	"github.com/labstack/echo"
)

func main() {
	router := echo.New()

	router.Get("/", func(c *echo.Context) error {
		ip := c.Request().Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = c.Request().RemoteAddr
		}
		// fmt.Println(ip)
		return c.JSON(http.StatusOK, ip)
	})

	// Start listening
	server := router.Server(":" + utils.GetEnvOrDefault("PORT", "3000"))

	server.TLSConfig = nil

	gracehttp.Serve(server)
}
