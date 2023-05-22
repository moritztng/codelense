package messaging

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
	Owner string
	Name  string
	Stars uint
}
