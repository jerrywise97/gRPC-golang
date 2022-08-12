package main

import (
	"coffee-shop/coffee"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("error dialing server::", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("error closing client conn::", err)

		}
	}(conn)
	coffeeService := coffee.NewCoffeeServiceClient(conn)
	response, err := coffeeService.AddCoffee(context.TODO(), &coffee.AddCoffeeRequest{
		Name:  "coffee black",
		Price: 30.00,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
