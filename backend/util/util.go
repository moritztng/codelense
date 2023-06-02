package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	Key               int
	Login             string
	Name              string
	Email             string
	Description       string
	MembersCount      int
	RepositoriesCount int
	TwitterUsername   string
	WebsiteUrl        string
	Url               string
	AvatarUrl         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type User struct {
	Key   int
	Login string
}

type Repository struct {
	Key        int
	Name       string
	StarsCount int
}

type Event struct {
	Key          int
	Type         string
	ActorId      int
	OrgId        int
	RepositoryId int
	Payload      json.RawMessage
	Public       bool
	CreatedAt    time.Time
}

func ReadConfig(configFile string) kafka.ConfigMap {

	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[parameter] = value
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}

	return m

}
