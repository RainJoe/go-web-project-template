package postgres

import "time"

//User defines the storage form of a user
type User struct {
	ID       int
	Name     string
	Email    string
	PassWord string
	Created  time.Time
}
