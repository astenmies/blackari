package model

// [TODO] change this to mongodb reviews
var Users = make(map[string][]*User)

type User struct {
	ID       string
	Username string `json:"username" bson:"username,omitempty" `
	Password string `json:"password" bson:"password,omitempty"`
}

type UserResolver struct {
	R *User
}

type UserInput struct {
	Username *string
	Password *string
}

type UserLoginInput struct {
	Username string
	Password string
}
