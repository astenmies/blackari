package db

import (
	"fmt"

	"github.com/astenmies/lychee/types"
)

type DB struct{}

func (p DB) Casual(s interface{ types.Sayer }) string {
	return fmt.Sprintf("Hey %s!", s.Say())
}
