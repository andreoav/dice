package transport

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/andreoav/dice/internal/app/parser"
	"github.com/andreoav/dice/pkg/api/dice"
)

// DiceService service
type DiceService struct {
	Parser parser.RollParser
}

// Roll endpoint
func (ds *DiceService) Roll(ctx context.Context, req *dice.RollRequest) (*dice.RollResponse, error) {
	fmt.Println("\nlets roll a:", req.GetInput())

	rollConfig, _ := ds.Parser.Parse(req.GetInput()) // TODO: handle err

	var rolls []int32

	for nRolls := int32(0); nRolls < rollConfig.Rolls; nRolls++ {
		roll := rand.Int31n(rollConfig.Dice) + rollConfig.Modifier + 1
		rolls = append(rolls, roll)
	}

	fmt.Println("rolls:", rolls)

	return &dice.RollResponse{
		Result: ds.sumValues(rolls),
	}, nil
}

func (DiceService) sumValues(values []int32) int32 {
	var total int32

	for _, value := range values {
		total += value
	}

	return total
}
