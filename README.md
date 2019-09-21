<p align="center">
  <strong>
    A simple example of GraphQL microservices federated in Golang, consumed by a Next.js client.
  </strong>
</p>

---
This repository is a work in progress...
---

## How to use
### Setup

```
go get -u github.com/astenmies/lychee
go get ./...
```

### Start

Open each microservice in a separate terminal tab. Start each of them with

```
cd micro-xxxxx
./start.sh
```

The federation of each of these microservices is done with the `federation` package.

```
cd federation
./start.sh
```


## Static assets

**lychee** uses [go-bindata](https://github.com/jteeuwen/go-bindata) to convert any file into managable Go source code. 
Useful for embedding binary data into a go program like GraphQL schemas.

### Usage

Write your microservice schema package, then generate `bindata.go`.

$ `go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...`

Here's how to use it from main.go.

```
// GetSchema returns the schema of Post
func GetSchema() string {
	s, _ := schema.Asset("schema/schema.graphql")
	stringSchema := string(s)

	return stringSchema
}
```