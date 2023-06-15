package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/model"
)

func main() {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	batchSize, err := strconv.Atoi(os.Getenv("BATCH_SIZE"))
	if err != nil {
		log.Fatal(err)
	}
	copyBuffer := make([][]any, batchSize)
	var gharchiveText string
	var gharchiveEvent model.GharchiveEvent
	for i := 0; ; i++ {
		rows, err := conn.Query(context.Background(), "select event from gharchive limit $1 offset $2", batchSize, i*batchSize)
		if err != nil {
			log.Fatal(err)
		}
		bufferIndex := 0
		_, err = pgx.ForEachRow(rows, []any{&gharchiveText}, func() error {
			json.Unmarshal([]byte(gharchiveText), &gharchiveEvent)
			payload, err := json.Marshal(gharchiveEvent.Payload)
			if err != nil {
				log.Fatal(err)
			}
			copyBuffer[bufferIndex] = []any{gharchiveEvent.ID, gharchiveEvent.Type, gharchiveEvent.Actor.ID, gharchiveEvent.Actor.Login, gharchiveEvent.Actor.DisplayLogin, gharchiveEvent.Actor.GravatarID, gharchiveEvent.Actor.URL, gharchiveEvent.Actor.AvatarURL, gharchiveEvent.Repo.ID, gharchiveEvent.Repo.Name, gharchiveEvent.Repo.URL, payload, gharchiveEvent.Public, gharchiveEvent.CreatedAt, gharchiveEvent.Org.ID, gharchiveEvent.Org.Login, gharchiveEvent.Org.GravatarID, gharchiveEvent.Org.URL, gharchiveEvent.Org.AvatarURL}
			bufferIndex++
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		nCopied, err := conn.CopyFrom(
			context.Background(),
			pgx.Identifier{"events"},
			[]string{"github_id", "type", "actor_id", "actor_login", "actor_display_login", "actor_gravatar_id", "actor_url", "actor_avatar_url", "repository_id", "repository_name", "repository_url", "payload", "public", "github_created_at", "org_id", "org_login", "org_gravatar_id", "org_url", "org_avatar_url"},
			pgx.CopyFromRows(copyBuffer[:bufferIndex]),
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(nCopied)
		if bufferIndex < batchSize {
			break
		}
	}
}
