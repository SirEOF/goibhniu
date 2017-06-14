package common

import "time"

type Reading struct {
	UpdateTime time.Time `json:"time"`
	Gravity    float64   `json:"gravity"`
}

func NewReading(update time.Time, gravity float64) Reading {
	return Reading{
		UpdateTime: update,
		Gravity:    gravity,
	}
}

func MakeReading(reading float64) Reading {
	return Reading{
		UpdateTime: time.Now(),
		Gravity:    reading,
	}
}
