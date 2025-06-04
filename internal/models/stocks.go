package models

import "time"

type Stock struct {
	ID         string    `db:"id"`
	Ticker     string    `db:"ticker"`
	Company    string    `db:"company"`
	TargetFrom float64   `db:"target_from"`
	TargetTo   float64   `db:"target_to"`
	RatingFrom string    `db:"rating_from"`
	RatingTo   string    `db:"rating_to"`
	Action     string    `db:"action"`
	Brokerage  string    `db:"brokerage"`
	Ts         time.Time `db:"ts"`
}
