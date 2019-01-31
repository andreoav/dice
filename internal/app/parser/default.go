package parser

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	// NumberOfRolls to run
	NumberOfRolls = iota + 1

	// DiceType to roll
	DiceType

	// Modifier of the roll
	Modifier
)

// var regex = regexp.MustCompile("(?i)([0-9]*)D([0-9]*)((?:[+/*-][0-9]+)|(?:[+-][LH]))?")
var regex = regexp.MustCompile("(?i)([0-9]*)D([0-9]*)([+/*-][0-9]+)?")

type defaultParser struct{}

func (dp defaultParser) Parse(input string) (*RollConfiguration, error) {
	matches := regex.FindStringSubmatch(input)

	return &RollConfiguration{
		Rolls:    dp.getRolls(matches),
		Dice:     dp.getDice(matches),
		Modifier: dp.getModifier(matches),
	}, nil
}

func (dp defaultParser) getRolls(matches []string) int32 {
	if NumberOfRolls >= len(matches) {
		return 1
	}

	return dp.getInt32Value(matches[NumberOfRolls], 1)
}

func (dp defaultParser) getDice(matches []string) int32 {
	if DiceType >= len(matches) {
		return 6 // TODO: add default dice
	}

	return dp.getInt32Value(matches[DiceType], 6)
}

func (dp defaultParser) getModifier(matches []string) int32 {
	if Modifier >= len(matches) {
		fmt.Println("whoops, mod")

		return 0
	}

	return dp.getInt32Value(matches[Modifier], 0)
}

func (defaultParser) getInt32Value(input string, defaultValue int32) int32 {
	value, err := strconv.ParseInt(input, 10, 32)

	if err != nil {
		return defaultValue
	}

	return int32(value)
}

// DefaultParser dice rolling parser
var DefaultParser = &defaultParser{}
