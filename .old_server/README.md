# lychee Server

## Initialize the schema

Install go-bindata
```bash
go get -u github.com/jteeuwen/go-bindata/...
```

Run the following command at root directory to generate Go code from .graphql file
```bash
# In lychee/server

go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
```

OR

```bash
# In lychee/server

go generate ./schema
```
There would be bindata.go generated under `schema` folder

## Start the server

```bash
# Get realize in order to watch changes in go apps
go get github.com/tockins/realize

# In lychee/server

# Start the server with realize
realize start --run --no-config
```

## Testing
```bash
# In lychee/server

go test ./test
```

## References
- https://github.com/graph-gophers/graphql-go
- https://github.com/tonyghita/graphql-go-example
- https://github.com/nilstgmd/graphql-starter-kit
- https://github.com/OscarYuen/go-graphql-starter
- https://blog.alexellis.io/golang-writing-unit-tests
- https://www.golang-book.com/books/intro/9
- http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/
- http://www.alexedwards.net/blog/golang-response-snippets