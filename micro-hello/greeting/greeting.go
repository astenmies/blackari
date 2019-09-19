package greeting

type GreetingResolver struct{}

// https://github.com/graph-gophers/graphql-go/issues/96#issuecomment-307663742
func (r *GreetingResolver) GetGreeting() string {
	return "Hello, world!"
}
