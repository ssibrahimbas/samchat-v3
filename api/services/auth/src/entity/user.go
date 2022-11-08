package entity

import "time"

type User struct {
	ID             string    `json:"uuid,omitempty" bson:"_id,omitempty"`
	Email          string    `json:"email" bson:"email"`
	Password       []byte    `bson:"password"`
	UserName       string    `json:"username,omitempty" bson:"username,omitempty"`
	FirstName      string    `json:"firstName,omitempty" bson:"first_name,omitempty"`
	LastName       string    `json:"lastName,omitempty" bson:"last_name,omitempty"`
	Avatar         string    `json:"avatar,omitempty" bson:"avatar,omitempty"`
	DateOfBirth    string    `json:"dateOfBirth,omitempty" bson:"date_of_birth,omitempty"`
	DateOfRegister string    `json:"dateOfRegister,omitempty" bson:"date_of_register,omitempty"`
	IsActive       bool      `json:"isActive,omitempty" bson:"is_active,omitempty"`
	IsCompleted    bool      `json:"isCompleted,omitempty" bson:"is_completed,omitempty"`
	LastIP         string    `json:"address,omitempty" bson:"last_ip,omitempty"`
	CreatedAt      time.Time `json:"dateOfCreate,omitempty" bson:"created_at"`
	UpdatedAt      time.Time `json:"dateOfUpdate,omitempty" bson:"updated_at,omitempty"`
}
