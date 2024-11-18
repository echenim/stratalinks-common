package models

import "time"

type Organization struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Country    string     `json:"country"`
	State      string     `json:"state"`
	City       string     `json:"city"`
	Created_At *time.Time `json:"created_at"`
	Updated_At *time.Time `json:"updated_at"`
}

type OrganizationBuilder struct {
	id        string
	name      string
	country   string
	state     string
	city      string
	createdAt *time.Time
	updatedAt *time.Time
}

func NewOrganizationBuilder() *OrganizationBuilder {
	return &OrganizationBuilder{}
}

// SetId sets the Id field of the organization
func (b *OrganizationBuilder) SetId(id string) *OrganizationBuilder {
	b.id = id
	return b
}

// SetName sets the Name field of the organization
func (b *OrganizationBuilder) SetName(name string) *OrganizationBuilder {
	b.name = name
	return b
}

// SetCountry sets the Country field of the organization
func (b *OrganizationBuilder) SetCountry(country string) *OrganizationBuilder {
	b.country = country
	return b
}

// SetState sets the State field of the organization
func (b *OrganizationBuilder) SetState(state string) *OrganizationBuilder {
	b.state = state
	return b
}

// SetCity sets the City field of the organization
func (b *OrganizationBuilder) SetCity(city string) *OrganizationBuilder {
	b.city = city
	return b
}

// SetCreatedAt sets the Created_At field of the organization
func (b *OrganizationBuilder) SetCreatedAt(createdAt *time.Time) *OrganizationBuilder {
	b.createdAt = createdAt
	return b
}

// SetUpdatedAt sets the Updated_At field of the organization
func (b *OrganizationBuilder) SetUpdatedAt(updatedAt *time.Time) *OrganizationBuilder {
	b.updatedAt = updatedAt
	return b
}

// Build constructs the Organization instance
func (b *OrganizationBuilder) Build() *Organization {
	return &Organization{
		Id:      b.id,
		Name:    b.name,
		Country: b.country,
		State:   b.state,
		City:    b.city,
	}
}
