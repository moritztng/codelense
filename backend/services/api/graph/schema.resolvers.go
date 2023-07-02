package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"log"
	"os"
	"strconv"

	pgx "github.com/jackc/pgx/v5"
	"github.com/moritztng/codelense/backend/services/api/graph/model"
)

// TimePoints is the resolver for the timePoints field.
func (r *queryResolver) TimePoints(ctx context.Context, fromDate int, toDate int, location string) ([]*model.TimePoint, error) {
	query, _ := os.ReadFile("query.sql")
	rows, err := r.Database.Query(context.Background(), string(query), fromDate, toDate, location)
	if err != nil {
		log.Fatal(err)
	}
	var time int
	var org_id int
	var stars_sum int
	timePoints := []*model.TimePoint{}
	timePoint := &model.TimePoint{}
	_, err = pgx.ForEachRow(rows, []any{&time, &org_id, &stars_sum}, func() error {
		if timePoint.Time != time {
			if timePoint.Time != 0 {
				timePoints = append(timePoints, timePoint)
			}
			timePoint = &model.TimePoint{Time: time}
		}
		timePoint.Values = append(timePoint.Values, &model.Value{Name: strconv.Itoa(org_id), Value: stars_sum})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return timePoints, nil
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, githubID int) (*model.Organization, error) {
	var login string
	var avatar_url string
	err := r.Database.QueryRow(context.Background(), "select login, avatar_url from organizations where github_id=$1", githubID).Scan(&login, &avatar_url)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Organization{Login: login, AvatarURL: avatar_url}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
