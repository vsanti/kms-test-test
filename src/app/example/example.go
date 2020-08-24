package example

import (
	"context"
	"fmt"

	"github.com/kyani-inc/kms-example/src/events"
	"github.com/kyani-inc/kms-example/src/services/log"
	pb "github.com/kyani-inc/proto/example"
	"go.opencensus.io/trace"
)

// SayHello says hello to the person
func SayHello(ctx context.Context, person *pb.Person) (string, error) {
	ctx, span := trace.StartSpan(ctx, "SayHello")
	defer span.End()

	log := log.Named("SayHello")
	log.Infof("%s said hello! Saying hi back!", person.GetName())

	// Build reply
	reply := fmt.Sprintf("Hello %s, nice to meet you!", person.GetName())

	// Emit event
	_, err := events.Worker.Publish(ctx, &pb.EventPersonSaidHello{
		Person: person,
		Reply:  reply,
	})

	return reply, err
}
