package db

import (
	"fmt"
	"testing"

	"github.com/KokoHere02/go-blog/config"
)

func TestDataBase(t *testing.T) {
	cfg := config.NewConfig()
	fmt.Print(cfg)
}
