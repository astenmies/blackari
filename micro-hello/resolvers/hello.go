package resolvers

type HelloResolver struct{}

// https://github.com/graph-gophers/graphql-go/issues/96#issuecomment-307663742
func (r *HelloResolver) SayHello() string {
	return "Hi there"
}
