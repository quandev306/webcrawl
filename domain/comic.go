package domain

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionComic = "comics"
)

type Comic struct {
	ID          primitive.ObjectID `bson:"_id"`
	SourceId   	string             `bson:"source_id"`
	Slug 	 	string             `bson:"slug"`
	Title      	string             `bson:"title"`
	Authors   	[]string           `bson:"authors"`
	Genres    	[]string           `bson:"genres"`
	Status     	string             `bson:"status"`
	Language   	string             `bson:"language"`
	CoverUrl  	string             `bson:"cover_url"`
	Description	string             `bson:"description"`
	UpdatedAt  	time.Time          `bson:"updated_at"`
	CreatedAt  	time.Time          `bson:"created_at"`
}

type ComicRepository interface {
	Create(c context.Context, comic *Comic) error
	Fetch(c context.Context) ([]Comic, error)
	GetBySourceIdAndSlug(c context.Context, sourceId string, slug string) (Comic, error)
}