package model

import "time"

type OptionsContract struct {
	Type          string    `json:"type"`
	StrikePrice   float64   `json:"strike_price"`
	Bid           float64   `json:"bid"`
	Ask           float64   `json:"ask"`
	ExpirationDate time.Time `json:"expiration_date"`
	LongShort     string    `json:"long_short"`
}
