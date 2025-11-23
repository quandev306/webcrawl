package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionJob = "jobs"
)

type Job struct {
	ID        primitive.ObjectID `bson:"_id"`
	Type      string             `bson:"type"`
	SourceId  string             `bson:"source_id"`
	ComicId   primitive.ObjectID `bson:"comic_id"`
	Status    string             `bson:"status"`
	Priority  int8               `bson:"priority"`
	Attempts  int8               `bson:"attempts"`
	NextRunAt time.Time		  	 `bson:"next_run_at"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type JobRepository interface {
	Create(c context.Context, job *Job) error
	Fetch(c context.Context) ([]Job, error)
}