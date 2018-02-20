# BLACKARI Server

## Initialize the schema

Install go-bindata
```bash
go get -u github.com/jteeuwen/go-bindata/...
```

Run the following command at root directory to generate Go code from .graphql file
```bash
# In blackari/server

go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go schema/...
```

OR

```bash
# In blackari/server

go generate ./schema
```
There would be bindata.go generated under `schema` folder

## Start the server

```bash
# Get realize in order to watch changes in go apps
go get github.com/tockins/realize

# In blackari/server

# Start the server with realize
realize start --run --no-config
```

## Testing
```bash
# In blackari/server

go test ./test
```

## References
- https://github.com/neelance/graphql-go
- https://github.com/tonyghita/graphql-go-example
- https://github.com/nilstgmd/graphql-starter-kit
- https://github.com/OscarYuen/go-graphql-starter
