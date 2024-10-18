package models

type Card struct {
	ID       int    `bson:"_id"`
	Name     string `bson:"name"`
	Type     string `bson:"type"`
	ImageURL string `bson:"image_url"`
}
