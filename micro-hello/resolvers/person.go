package resolvers

type PersonResolver struct {
	name string `json:"name"`
}

func (r *PersonResolver) To() (*PersonResolver, error) {
	s := PersonResolver{
		name: "Sebastien",
	}

	return &s, nil

}

// Title resolves the title field for Post
func (p *PersonResolver) Name() *string {
	return &p.name
}
