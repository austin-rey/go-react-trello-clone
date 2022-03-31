package organization

type Organization struct {
	OrganizationID   int    `json:"organizationId"`
	OrganizationName string `json:"organizationName"`
	OrganizationDesc string `json:"organizationDesc"`
}