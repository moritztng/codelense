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

type Event struct {
	GithubId        uint            `json:"github_id"`
	Type            string          `json:"type"`
	ActorId         uint            `json:"actor_id"`
	ActorLogin      string          `json:"actor_login"`
	ActorUrl        string          `json:"actor_url"`
	ActorAvatarUrl  string          `json:"actor_avatar_url"`
	RepositoryId    uint            `json:"repository_id"`
	RepositoryName  string          `json:"repository_name"`
	RepositoryUrl   string          `json:"repository_url"`
	Payload         json.RawMessage `json:"payload"`
	Public          bool            `json:"public"`
	GithubCreatedAt time.Time       `json:"github_created_at"`
	OrgId           uint            `json:"org_id"`
	OrgLogin        string          `json:"org_login"`
	OrgUrl          string          `json:"org_url"`
	OrgAvatarUrl    string          `json:"org_avatar_url"`
}
