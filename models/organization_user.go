package models

import "time"

type User struct {
	Id             string     `json:"id"`
	OrgId          string     `json:"organization_id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Email          string     `json:"email"`
	EmailConfirmed bool       `json:"email_confirmed"`
	Status         string     `json:"status"`
	Created_At     time.Time  `json:"created_at"`
	Updated_At     *time.Time `json:"updated_at"`
}

// UserBuilder defines the builder structure
type UserBuilder struct {
	id        string
	orgId     string
	firstName string
	lastName  string
	email     string
	status    string
}

// NewUserBuilder creates a new instance of UserBuilder
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) SetId(id string) *UserBuilder {
	b.id = id
	return b
}

func (b *UserBuilder) SetOrgId(orgId string) *UserBuilder {
	b.orgId = orgId
	return b
}

func (b *UserBuilder) SetFirstName(firstName string) *UserBuilder {
	b.firstName = firstName
	return b
}

func (b *UserBuilder) SetLastName(lastName string) *UserBuilder {
	b.lastName = lastName
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.email = email
	return b
}

func (b *UserBuilder) SetStatus(status string) *UserBuilder {
	b.status = status
	return b
}

// Build constructs the final User instance
func (b *UserBuilder) Build() User {
	return User{
		Id:        b.id,
		OrgId:     b.orgId,
		FirstName: b.firstName,
		LastName:  b.lastName,
		Email:     b.email,
		Status:    b.status,
	}
}
