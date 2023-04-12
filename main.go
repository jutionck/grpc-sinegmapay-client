package main

import (
	"context"
	"github.com/jutionck/grpc-sinegmapay-client/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	srv := service.NewSinegmaPaymentClient(conn)
	response, err := srv.CheckBalance(context.Background(), &service.CheckBalanceMessage{CustomerId: 3})
	if err != nil {
		log.Fatalf("error when calling check balance: %v", err)
	}
	log.Println("Balance:", response)

	response, err = srv.DoPayment(context.Background(), &service.PaymentMessage{
		CustomerId: 4,
		Amount:     35000,
	})
	log.Println("Balance:", response)
}
