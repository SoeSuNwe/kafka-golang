package main

import (
	"context"
	"flag"
	"log"
	pbEvent "map/cmd/vas-integration-apiserver/grpc/AnalyticsEvent"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbEvent.NewAnalyticsEventServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	r, err := c.SendEvent(ctx, &pbEvent.AnalyticsEvent{
		TaskId:    "taskId is one",
		EngineId:  "enginId1",
		SourceId:  "s1",
		TimeStamp: "",
		EventId:   "ev1",
		Snapshot:  []byte("SnapshotSample"),
		EventData: []byte("EventDataSample"),
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
