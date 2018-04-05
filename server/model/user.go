package model

// [TODO] change this to mongodb reviews
var Users = make(map[string][]*User)

type User struct {
	Username *string `json:"username" bson:"username,omitempty" `
	Password *string `json:"password" bson:"password,omitempty"`
}

type UserResolver struct {
	R *User
}

type UserInput struct {
	Username *string
	Password *string
}
