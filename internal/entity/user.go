// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// User -.
type User struct {
	ID        int    `json:"id"         example:"1"`
	Username  string `json:"username"   example:"rmscoal"`
	FirstName string `json:"firstName"  example:"John"`
	LastName  string `json:"lastName"   example:"Smith"`
	Password  string `json:"password"   example:"fnjsfkanfja"`
}
