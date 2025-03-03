package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeoJSON struct {
	Type        string      `bson:"type" json:"type"`
	Coordinates interface{} `bson:"coordinates" json:"coordinates"`
}

// Struct vùng đất
type Land struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Location GeoJSON            `bson:"location" json:"location"`
	Type     string             `bson:"type" json:"type"`
	Area     float64            `bson:"area" json:"area"`
}
