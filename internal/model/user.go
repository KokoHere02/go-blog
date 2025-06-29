package model

import "time"

type User struct {
	id         int
	username   string
	phone      string
	email      string
	password   string
	nickname   string
	avatar_url string
	bio        string
	is_admin   bool
	created_at time.Time
	updated_at time.Time
}
