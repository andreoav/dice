package parser

// RollConfiguration struct
type RollConfiguration struct {
	Rolls    int32
	Dice     int32
	Modifier int32
}

// RollParser to implement
type RollParser interface {
	Parse(input string) (*RollConfiguration, error)
}
