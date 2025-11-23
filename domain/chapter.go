package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionChapter = "chapters"
)

type Chapter struct {
	ID              primitive.ObjectID `bson:"_id"`
	ComicId         primitive.ObjectID `bson:"comic_id"`
	SourceChapterId string             `bson:"source_chapter_id"`
	Name            string             `bson:"name"`
	Number          float64            `bson:"number"`
	Volume          float64            `bson:"volume"`
	Pages           []Page             `bson:"pages"`
	PublishedAt     time.Time          `bson:"published_at"`
	Language        string             `bson:"language"`
	CreatedAt       time.Time          `bson:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at"`
}

type ChapterRepository interface {
	Create(c context.Context, chapter *Chapter) error
	Fetch(c context.Context) ([]Chapter, error)
	GetByComicIdAndSourceChapterId(c context.Context, comicId primitive.ObjectID, sourceChapterId string) (Chapter, error)
}
