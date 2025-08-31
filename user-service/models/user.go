package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Userid    primitive.ObjectID `bson:"_id,omitempty" json:"userid"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	Role      string             `json:"role" bson:"role"`
	OrgName   string             `json:"orgName" bson:"orgName"`
	ContactNo string             `json:"contactNo" bson:"contactNo"`
	Address   string             `json:"address" bson:"address"`
}
