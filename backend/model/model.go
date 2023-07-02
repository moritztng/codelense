package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	GithubId        uint      `json:"github_id"`
	Type            string    `json:"type"`
	ActorId         uint      `json:"actor_id"`
	ActorLogin      string    `json:"actor_login"`
	ActorUrl        string    `json:"actor_url"`
	ActorAvatarUrl  string    `json:"actor_avatar_url"`
	RepositoryId    uint      `json:"repository_id"`
	RepositoryName  string    `json:"repository_name"`
	RepositoryUrl   string    `json:"repository_url"`
	Public          bool      `json:"public"`
	GithubCreatedAt time.Time `json:"github_created_at"`
	OrgId           uint      `json:"org_id"`
	OrgLogin        string    `json:"org_login"`
	OrgUrl          string    `json:"org_url"`
	OrgAvatarUrl    string    `json:"org_avatar_url"`
}

type Organization struct {
	gorm.Model
	GithubID        uint `gorm:"unique"`
	Login           string
	Name            string
	Email           string
	Description     string
	Location        string
	TwitterUsername string
	WebsiteUrl      string
	Url             string
	AvatarUrl       string
	GithubCreatedAt time.Time
	GithubUpdatedAt time.Time
}
