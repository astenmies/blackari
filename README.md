# Lychee
👻 A Go GraphQL server with a React SSR client (Next.js)

WORK IN PROGRESS

## Getting started

```bash
# Clone this repository
git clone https://github.com/astenmies/lychee.git

# Change directory
cd lychee

# Rename .env.example to .env and fill it with your configs
cp client/.env.example client/.env
```

## [Server docs](./server)

## [Client docs](./client)


## Troubleshooting

### Mongodb returns an empty object
This means that your fields are not exported!

You have something like this:

```
type Post struct {
	id    uint32
	title string `json:"title"`
}

func GetPost() {
    var result Post
	collection := db.Client.Database("lychee").Collection("post")
	err := collection.FindOne(context.TODO(), bson.M{"id": uint32(2)}).Decode(&result)
}
```

Replace by this:

```
type Post struct {
	ID    uint32
	Title string `json:"title"`
}
```

### Cannot query field "node"
https://github.com/nautilus/gateway/issues/74#issuecomment-519987797