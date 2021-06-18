package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model for mongo DB
type User struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name string `bson:"name" json:"name,omitempty"` 
	Surname string `bson:"surname" json:"surname,omitempty"` 
	BirthDate time.Time `bson:"birthDate" json:"birthDate,omitempty"` 
	Email string `bson:"email" json:"email,omitempty"` 
	Password string `bson:"password" json:"password,omitempty"` 
	Avatar string `bson:"avatar" json:"avatar,omitempty"` 
	Banner string `bson:"banner" json:"banner,omitempty"` 
	Biography string `bson:"biography" json:"biography,omitempty"` 
	Location string `bson:"location" json:"location,omitempty"` 
	Website string `bson:"website" json:"website,omitempty"` 
}
