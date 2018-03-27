# lychee Client

## Start the client

Make sure the server is running before generating the schema.

```bash
# In lychee/client

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
# In lychee/client

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