package main

import (
	"time"
)

type ID = int // user id

type MatchRequest struct {
	UserID    ID
	Name      string
	Gender    Gender
	Height    int
	Dates     int
	CreatedAt time.Time
}

type Gender string

const (
	GENDER_KNOWN  Gender = "UNKNOWN"
	GENDER_MALE   Gender = "MALE"
	GENDER_FEMALE Gender = "FEMALE"
)
