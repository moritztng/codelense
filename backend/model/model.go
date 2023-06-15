package model

import (
	"encoding/json"
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

type GharchiveEventJson struct {
	Event datatypes.JSON
}

type GharchiveEvent struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		ID           int    `json:"id"`
		Login        string `json:"login"`
		DisplayLogin string `json:"display_login"`
		GravatarID   string `json:"gravatar_id"`
		URL          string `json:"url"`
		AvatarURL    string `json:"avatar_url"`
	} `json:"actor"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	Public    bool            `json:"public"`
	CreatedAt time.Time       `json:"created_at"`
	Org       struct {
		ID         int    `json:"id"`
		Login      string `json:"login"`
		GravatarID string `json:"gravatar_id"`
		URL        string `json:"url"`
		AvatarURL  string `json:"avatar_url"`
	} `json:"org"`
}
