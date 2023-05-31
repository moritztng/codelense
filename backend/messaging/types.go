package messaging

import (
	"encoding/json"
	"time"
)

type ApiGithubRequest struct {
	Key      string
	MaxStars uint
	First    uint
}

type GithubStoreResult struct {
	Key          string
	Repositories []*Organization
}

type Organization struct {
	Key       int
	Login     string
	Name      string
	CreatedAt time.Time
}

type Event struct {
	Key       int
	Type      string
	Payload   json.RawMessage
	CreatedAt time.Time
}
