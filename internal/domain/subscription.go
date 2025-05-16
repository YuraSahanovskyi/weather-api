package domain

import (
	"fmt"
	"strings"
)

type Subscription struct {
	Email     string
	City      string
	Frequency Frequency
	Confirmed bool
	Token string
}

type Frequency int

const (
	Hourly Frequency = iota
	Daily
)

var frequencyName = map[Frequency]string{
	Hourly: "hourly",
	Daily:  "daily",
}

var frequencyValue = map[string]Frequency{
	"hourly": Hourly,
	"daily":  Daily,
}

func (f Frequency) String() string {
	return frequencyName[f]
}

func ParseFrequency(frequencyName string) (*Frequency, error) {
	if frequency, ok := frequencyValue[strings.ToLower(frequencyName)]; ok {
		return &frequency, nil
	}
	return nil, fmt.Errorf("invalid frequency: %s", frequencyName)
}
