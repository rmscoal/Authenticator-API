// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "database/sql"

// User -.
type User struct {
	ID        int            `json:"id"         example:"1"`
	Username  string         `json:"username"   example:"rmscoal"`
	Email     string         `json:"email" example:"hello@gmail.com"`
	FirstName sql.NullString `json:"firstName"  example:"John"`
	LastName  sql.NullString `json:"lastName"   example:"Smith"`
	Password  string         `json:"password"   example:"fnjsfkanfja"`
}
