package db

import (
	"testing"
	"time"
)

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

func TestDataBase(t *testing.T) {
	// cfg := config.NewConfig()
	// var user User
	// db := DB.Find(&user)
	// if db.Error != nil {
	// 	t.Errorf("Failed to find user: %v", db.Error)
	// }
	// fmt.Println(user)
}
