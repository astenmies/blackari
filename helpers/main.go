package helpers

import (
	"fmt"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

func IDToUint(i graphql.ID) (uint, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, errors.Wrap(err, "GqlIDToUint")
	}

	return uint(r), nil
}

func Int32P(i uint) *int32 {
	r := int32(i)
	return &r
}

func BoolP(b bool) *bool {
	return &b
}

func IDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}
