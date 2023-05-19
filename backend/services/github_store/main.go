package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/moritztng/codelense/backend/messaging"
)

type Repo struct {
	Name  string
	Stars uint
}

func main() {
	//var text string
	conn, _ := pgx.Connect(context.Background(), "postgresql://postgres:example@localhost:5001/postgres")
	defer conn.Close(context.Background())
	topic := "github"
	conf := messaging.ReadConfig("kafka.properties")
	conf["group.id"] = "github_store"
	conf["auto.offset.reset"] = "earliest"
	c, _ := kafka.NewConsumer(&conf)
	c.SubscribeTopics([]string{topic}, nil)
	run := true
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err != nil {
			continue
		}
		var repo Repo
		json.Unmarshal(msg.Value, &repo)
		tx, _ := conn.Begin(context.Background())
		defer tx.Rollback(context.Background())
		_, err = tx.Exec(context.Background(), "insert into github(name, stars) values ($1,$2)", repo.Name, repo.Stars)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = tx.Commit(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	c.Close()
}
