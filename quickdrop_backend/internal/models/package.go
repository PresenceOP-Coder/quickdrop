package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Package struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Address   string             `bson:"address" json:"address"`
	Latitude  float64            `bson:"latitude" json:"latitute"`
	Longitude float64            `bson:"longitude" json:"longitude"`
	Status    string             `bson:"status" json:"status"`
}

type Driver struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name"`
	CurrentLat float64            `bson:"current_lat" json:"current_lat"`
	CurrentLng float64            `bson:"current_lng" json:"current_lng"`
	Status     string             `bson:"status" json:"status"`
	ActivePackages []primitive.ObjectID `bson:"active_packages" json:"active_packages"`
}
