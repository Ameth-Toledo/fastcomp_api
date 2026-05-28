package entities

import "time"

type User struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	BusinessName     string    `json:"businessName"`
	Email            string    `json:"email"`
	Password         *string   `json:"-"`
	Phone            string    `json:"phone"`
	Website          *string   `json:"website,omitempty"`
	ProfilePhoto     *string   `json:"profilePhoto,omitempty"`
	RegistrationDate time.Time `json:"registrationDate"`
	RoleID           int       `json:"roleId"`
}
