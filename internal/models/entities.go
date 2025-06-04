package models

import "time"

type Company struct {
	ID     int64  `db:"company_id"`
	Ticker string `db:"ticker"`
	Name   string `db:"name"`
}

type Brokerage struct {
	ID   int64  `db:"brokerage_id"`
	Name string `db:"name"`
}

type RatingType struct {
	ID          int64  `db:"rating_id"`
	Code        string `db:"code"`
	Description string `db:"description"`
}

type ActionType struct {
	ID          int64  `db:"action_id"`
	Code        string `db:"code"`
	Description string `db:"description"`
}

type AnalystUpdate struct {
	ID           int64     `db:"update_id"`
	CompanyID    int64     `db:"company_id"`
	BrokerageID  int64     `db:"brokerage_id"`
	ActionID     int64     `db:"action_id"`
	RatingFromID int64     `db:"rating_from_id"`
	RatingToID   int64     `db:"rating_to_id"`
	TargetFrom   float64   `db:"target_from"`
	TargetTo     float64   `db:"target_to"`
	EventTime    time.Time `db:"event_time"`
}
