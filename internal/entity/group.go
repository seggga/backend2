package entity

import "github.com/google/uuid"

// Group represents a group of users (project, organization, public)
type Group struct {
	ID          uuid.UUID
	Type        string
	Name        string
	Description string
}
