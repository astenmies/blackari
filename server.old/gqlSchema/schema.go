//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
package gqlSchema

import (
	"bytes"

	"github.com/graph-gophers/graphql-go"
)

var GraphqlSchema *graphql.Schema

// GetRootSchema describes the data that we ask for
func GetRootSchema() string {

	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
