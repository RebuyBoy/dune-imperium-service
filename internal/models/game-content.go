package models

type GameContent struct {
	ID       string      `bson:"_id"`
	Name     string      `bson:"name"`
	Type     ContentType `bson:"type"`
	ImageURL string      `bson:"image_url"`
}

type ContentType string

const (
	Leader   ContentType = "leader"
	Tech     ContentType = "tech"
	Intrigue ContentType = "intrigue"
	Card     ContentType = "card"
)
