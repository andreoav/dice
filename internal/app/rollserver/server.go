package rollserver

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/andreoav/dice/internal/app/dice/transport"
	"github.com/andreoav/dice/internal/app/parser"
	"github.com/andreoav/dice/pkg/api/dice"
	"google.golang.org/grpc"
)

// Configuration options
type Configuration struct {
	Address string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Start a gRPC server
func Start(config *Configuration) error {
	fmt.Println("Starting Roll gRPC Server...")

	listen, err := net.Listen("tcp", config.Address)

	if err != nil {
		return err
	}

	s := grpc.NewServer()

	dice.RegisterRollServiceServer(s, &transport.DiceService{
		Parser: parser.DefaultParser,
	})

	return s.Serve(listen)
}
