package main

import (
	"os"

	"github.com/kyani-inc/kms-example/src/events"
	"github.com/kyani-inc/kms-example/src/rest"
	"github.com/kyani-inc/kms-example/src/rpc"
	"github.com/kyani-inc/kms-example/src/services/log"
	"github.com/kyani-inc/kms/v2"
)

var (
	// AppName from build
	AppName string

	// BUILD number
	BUILD string
)

func main() {
	// Copy build info to environment
	os.Setenv("AppName", AppName)
	os.Setenv("BUILD", BUILD)

	// Create KMS service & bind handlers
	server := kms.NewService(kms.ServiceName("example"))
	server.EnableRPC(rpc.Setup)
	server.EnableREST(rest.Setup)
	server.EnableWorker(events.Setup)

	// Alias the server log instance
	log.Register(server.Logger())

	// Do our app-specific setup/initialization here

	// Setup database (requires credentials)
	// database.Setup()

	// Start the server
	server.Start()
}
