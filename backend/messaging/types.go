package messaging

import "time"

type ApiGithubRequest struct {
	Key      string
	MaxStars uint
	First    uint
}

type GithubStoreResult struct {
	Key          string
	Repositories []*Repository
}

type Repository struct {
	Key       int
	Login     string
	Name      string
	CreatedAt time.Time
}
