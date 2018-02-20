# Blackari
👻 A Go GraphQL server with a React SSR client (Next.js)

WORK IN PROGRESS

## Getting started

```bash
# Clone this repository
git clone https://github.com/astenmies/blackari.git

# Change directory
cd blackari

# Rename .env.example to .env and fill it with your configs
mv client/.env.example client/.env
```

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

# Start the server with realize
realize start --run --no-config
```

## Start the client

Make sure the server is running before generating the schema.

```bash
# In blackari/client

# Install node dependencies
yarn

# Run relay
yarn run relay

# Generate the GraphQL schema
yarn run schema

# Run the app
yarn run dev
```

## .env

Here are a few examples of .env usage.

https://github.com/zeit/next.js/tree/9320d9f006164f2e997982ce76e80122049ccb3c/examples/with-dotenv

(not as good as previous example) https://github.com/iaincollins/nextjs-starter/blob/master/next.config.js


## Troubleshooting



### SyntaxError: Unexpected token import
This is almost always related to babel presets.
```bash
yarn add babel-preset-es2015 -D
```
Then add it in **.babelrc**. This is an example of a full minimal babel config file.
```
{
  "presets": [
    "es2015",
    "stage-1",
    "react"
  ],
  "plugins": [
    "transform-object-assign"
  ]
}