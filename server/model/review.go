package model

// [TODO] change this to mongodb reviews
var Reviews = make(map[string][]*Review)

type Review struct {
	// ID         string  `json:"id" bson:"_id,omitempty"`
	Stars      *int32  `json:"stars" bson:"stars,omitempty" `
	Commentary *string `json:"commentary" bson:"commentary,omitempty"`
}

type ReviewResolver struct {
	R *Review
}

type ReviewInput struct {
	Stars      *int32
	Commentary *string
}
