package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Eklund2012/recall/proto/gen/study_service"
)

func CallSearch(query string) {
	// 1. Establish connection to Python
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStudyEngineClient(conn)

	// 2. Prepare the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 3. Make the call
	r, err := c.SearchNotes(ctx, &pb.SearchRequest{Query: query, TopK: 3})
	if err != nil {
		log.Fatalf("could not search: %v", err)
	}

	// 4. Print results
	for _, match := range r.GetMatches() {
		log.Printf("Match: %s (Source: %s)", match.GetText(), match.GetSource())
	}
}
