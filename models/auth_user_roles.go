package models

type UserAssignedRoles struct {
	UserId string `json:"user"`
	RoleId string `json:"role_id"`
}
