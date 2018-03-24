package utils

import "reflect"

func AllNil(x interface{}) bool {
	rv := reflect.ValueOf(x)
	rv = rv.Elem()

	for i := 0; i < rv.NumField(); i++ {
		if !rv.Field(i).IsNil() {
			return false
		}
	}
	return true
}

// Check if an interface is nil vs. check if interface field is nil
// https://stackoverflow.com/questions/38070015/check-if-interface-value-is-nil-in-go-without-using-reflect
// There are two things: If y is the nil interface itself (in which case y==nil will be true)
// or if y is a non-nil interface but underlying value is a nil value (in which case y==nil will be false).
// https://play.golang.org/p/BjDd1ko94E
