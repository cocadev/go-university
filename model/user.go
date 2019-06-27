package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// omit is the bool type for omitting a field of struct.
type omit bool

// Verify struct.
type Verify struct {
	IsVerified bool      `json:"isVerified" bson:"isVerified"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
}

// User is a user model
type User struct {
	ID                bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty" description:"Object ID"`
	Avatar            string        `json:"avatar" description:"User avatar url"`
	Firstname         string        `json:"firstname" description:"User firstname"`
	Lastname          string        `json:"lastname" description:"User lastname"`
	FacebookUserID    string        `json:"facebookUserId" bson:"facebookUserId" description:"When user login or register signup with Facebook, this field will be updated"`
	Email             string        `json:"email" description:"User email address"`
	Birth             time.Time     `json:"birth" bson:"birth,omitempty" description:"User birth"`
	Password          string        `json:"password"`
	CountryCode       string        `json:"countryCode" bson:"countryCode"`
	PhoneCode         string        `json:"phoneCode" bson:"phoneCode"`
	Phone             string        `json:"phone"`
	Verify            *Verify       `json:"verify,omitempty" bson:"verify,omitempty"`
	OneSignalPlayerID string        `json:"onesignalPlayerId,omitempty" bson:"onesignalPlayerId,omitempty"`
	CreatedAt         time.Time     `json:"createdAt" bson:"createdAt" description:"User created date"`
	UpdatedAt         time.Time     `json:"updatedAt" bson:"updatedAt" description:"User updated date. This field will be updated when any update operation will be occured"`
}

// PublicUser struct.
type PublicUser struct {
	*User
	Password          omit `json:"password,omitempty"`
	Verify            omit `json:"verify,omitempty"`
	OneSignalPlayerID omit `json:"onesignalPlayerId,omitempty"`
}

// Fullname method
func (user *User) Fullname() string {
	return user.Firstname + " " + user.Lastname
}
