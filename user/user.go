package user

type User struct {
	UserID         int    `json:"userId"`
	OrganizationID int    `json:"organizationId"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	JobTitle       string `json:"jobTitle"`
}