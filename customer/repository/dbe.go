package repository

import "time"

type CreateCustomerDbe struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name,omitempty"`
	Email     string    `bson:"email,omitempty"`
	Phone     string    `bson:"phone,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}

type CustomerDbe struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name,omitempty"`
	Email     string    `bson:"email,omitempty"`
	Phone     string    `bson:"phone,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
}
