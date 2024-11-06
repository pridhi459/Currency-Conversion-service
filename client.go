package main

import (
	"context"
	"fmt"
	"log"

	pb "currency-conversion-service/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCurrencyConverterClient(conn)

	// Example conversion request
	req := &pb.CurrencyRequest{
		From:   "USD",
		To:     "EUR",
		Amount: 100,
	}

	response, err := client.Convert(context.Background(), req)
	if err != nil {
		log.Fatalf("could not convert: %v", err)
	}

	fmt.Printf("Converted %.2f %s to %.2f %s\n", response.Amount, response.From, response.Converted, response.To)
}
