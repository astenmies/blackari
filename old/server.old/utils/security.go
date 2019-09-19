package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt :
// - Uses GenerateFromPassword to hash & salt pwd.
// - MinCost is just an integer constant provided by the bcrypt
// package along with DefaultCost & MaxCost.
// - The cost can be any value you want provided it isn't lower
// than the MinCost (4)-
func HashAndSalt(pwd string) string {
	// use *pwd to de-reference the pointer
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
