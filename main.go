package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "currency-conversion-service/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCurrencyConverterServer
}

// Hardcoded exchange rates
var exchangeRates = map[string]float64{
	"USD": 1.0, // Base currency
	"EUR": 84.0,
	"INR": 0.92,
}

// Convert performs the currency conversion
func (s *server) Convert(ctx context.Context, req *pb.CurrencyRequest) (*pb.CurrencyResponse, error) {

	fromRate, fromExists := exchangeRates[req.From]
	toRate, toExists := exchangeRates[req.To]

	if !fromExists || !toExists {
		return nil, fmt.Errorf("unsupported currency")
	}

	// Convert to base currency (USD)
	amountInUSD := float64(req.Amount) / fromRate
	// Convert from base currency to target currency
	convertedAmount := amountInUSD * toRate

	return &pb.CurrencyResponse{
		From:      req.From,
		To:        req.To,
		Amount:    req.Amount,
		Converted: float32(convertedAmount),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(grpcServer, &server{})

	fmt.Println("Currency Conversion Service is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
