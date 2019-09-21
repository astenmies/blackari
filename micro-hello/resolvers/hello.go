package resolvers

type HelloResolver struct{}

// https://github.com/graph-gophers/graphql-go/issues/96#issuecomment-307663742
func (r *HelloResolver) SayHello() string {
	return "Hi there"
}

// type PostResolver struct {
// 	DB    *db.Services
// 	model models.Post
// }

// func (r *PostResolver) Subtitle() string {
// 	return "Hooooola"

// }

// func (q *PostResolver) Node(args struct{ ID string }) *types.NodeResolver {
// 	// user := users[args.ID]

// 	return &types.NodeResolver{r.model}
// }
