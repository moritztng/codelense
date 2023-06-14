package main

import (
	"bufio"
	"compress/gzip"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	datetime, _ := time.Parse("2006-01-02", os.Getenv("START_DATE"))
	copyBufferLength, _ := strconv.Atoi(os.Getenv("COPY_BUFFER_LENGTH"))
	for ; datetime.Before(time.Now()); datetime = datetime.Add(time.Hour) {
		url := fmt.Sprintf("https://data.gharchive.org/%d-%02d-%02d-%d.json.gz", datetime.Year(), datetime.Month(), datetime.Day(), datetime.Hour())
		log.Println(url)
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		gzipReader, err := gzip.NewReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(gzipReader)
		scannerBuffer := make([]byte, 0, 64*1024)
		scanner.Buffer(scannerBuffer, 1024*1024)
		copyBuffer := make([][]any, copyBufferLength)
	loop:
		for {
			for rowIndex := 0; rowIndex < copyBufferLength; rowIndex++ {
				if !scanner.Scan() {
					if scanner.Err() != nil {
						log.Fatal(scanner.Err())
					}
					break loop
				}
				copyBuffer[rowIndex] = []any{scanner.Text()}
			}
			nCopied, err := conn.CopyFrom(
				context.Background(),
				pgx.Identifier{"gharchive"},
				[]string{"event"},
				pgx.CopyFromRows(copyBuffer),
			)
			if err != nil {
				log.Fatal(nCopied, err)
			}
			log.Println(nCopied)
		}
	}
}
