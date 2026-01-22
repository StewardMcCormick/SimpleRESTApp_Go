package model

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}
