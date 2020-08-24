package rest

import (
	"net/http"

	"github.com/kyani-inc/kms-example/src/app/example"
	"github.com/kyani-inc/kms-example/src/services/log"
	pb "github.com/kyani-inc/proto/example"
	"github.com/labstack/echo/v4"
)

// Setup configures the service's Echo (REST) server
func Setup(server *echo.Echo) {
	server.GET("/hello/:name", sayHello)
}

// sayHello is the route for saying hello
func sayHello(c echo.Context) error {
	log := log.Named("sayHello")

	// Grab name from query string
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "missing name parameter",
		})
	}

	// Build our person using Proto type
	person := &pb.Person{
		Name: name,
	}

	// Say hello using the app package
	answer, err := example.SayHello(c.Request().Context(), person)
	if err != nil {
		log.Errorf("error saying hello: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	// Return it
	return c.JSON(http.StatusOK, map[string]string{
		"reply": answer,
	})
}
