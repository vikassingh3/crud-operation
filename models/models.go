package models

import "time"

type User struct {
	ID          *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        *string   `json:"name"`
	Designation *string   `json:"designation"`
	Age         *int      `json:"age"`
	Created_At  time.Time `json:"created_at"`
}
