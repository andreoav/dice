package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/andreoav/dice/pkg/api/dice"

	"google.golang.org/grpc"
)

func main() {
	roll := flag.String("dice", "1d6", "dice config to roll")

	flag.Parse()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("error connecting to grpc server:", err)
	}

	defer conn.Close()

	client := dice.NewRollServiceClient(conn)

	res, err := client.Roll(context.Background(), &dice.RollRequest{
		Input: *roll,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Response: ", res.Result)
}
