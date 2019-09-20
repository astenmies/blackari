package db

import (
	"fmt"
)

type DB struct {
	DB string
}

func (p DB) Casual(s string) string {
	return fmt.Sprintf("Hey %s!", s)
}
