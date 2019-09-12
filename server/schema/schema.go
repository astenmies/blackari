//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
package schema

import (
	"bytes"

	"github.com/davecgh/go-spew/spew"
	graphql "github.com/graph-gophers/graphql-go"
)

// Points to graphql schema
var GraphqlSchema *graphql.Schema

// GetRootSchema describes the data that we ask for
func GetRootSchema() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		spew.Dump("name", name)
		b := MustAsset(name)
		buf.Write(b)

		// Add a newline if the file does not end with a newline
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
