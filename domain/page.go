package domain

const (
	CollectionPage = "pages"
)

type Page struct {
	Number 	 	int16            `bson:"number"`
	URL 	 	string           `bson:"url"`
	Hash 	 	string           `bson:"hash,omitempty"`
}
