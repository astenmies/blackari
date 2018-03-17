package utils

import (
	"log"
	"os"
)

func Exists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		log.Fatalf(err.Error())
	}

	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
