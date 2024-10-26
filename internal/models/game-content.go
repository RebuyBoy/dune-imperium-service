package models

type GameContent struct {
	ID        string      `bson:"_id" json:"id"`
	Name      string      `bson:"name" json:"name"`
	Type      ContentType `bson:"type" json:"type"`
	ImageURL  string      `bson:"image_url" json:"url"`
	Expansion Expansion   `bson:"expansion" json:"expansion"`
}

type ContentType string
type Expansion string

const (
	Leader   ContentType = "leader"
	Tech     ContentType = "tech"
	Intrigue ContentType = "intrigue"
	Card     ContentType = "card"
)

const (
	BaseGame Expansion = "base"
	RiseOfIx Expansion = "rise_of_ix"
)
