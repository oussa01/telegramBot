package models

import "time"

type User struct {
	ID               int
	Answered         map[string]bool
	XP               int
	Attempts         int
	DailyCount       int
	LastDay          string
	SelectingMode    bool
	CurrentChallenge *Challenge
}

func (u *User) CanPlayToday() bool {
	today := time.Now().Format("2006-01-02")
	if u.LastDay != today {
		u.DailyCount = 0
		u.LastDay = today
	}
	return u.DailyCount < 3
}

func (u *User) MarkPlayed() {
	u.DailyCount++
}
