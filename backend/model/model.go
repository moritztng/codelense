package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Organization struct {
	GithubID        uint
	Login           string
	Name            string
	Email           string
	Description     string
	TwitterUsername string
	WebsiteUrl      string
	Url             string
	AvatarUrl       string
	GithubCreatedAt time.Time
	GithubUpdatedAt time.Time
}

type OrganizationEvent struct {
	gorm.Model
	Organization Organization `gorm:"embedded;embeddedPrefix:organization_"`
}

type Event struct {
	gorm.Model
	GithubID        uint
	Type            string
	ActorID         uint
	OrgID           uint
	RepositoryID    uint
	Payload         datatypes.JSON
	Public          bool
	GithubCreatedAt time.Time
}
