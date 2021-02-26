package entity

// Raindrop defines the data type to save a bookmark.
// It is named Raindrop because it is integrated with raindrop.io
type Raindrop struct {
	// CollectionName defines in which collection the bookmark is saved.
	CollectionName string `json:"collection_name"`
	// Link defines the link of the bookmark.
	Link string `json:"link"`
}
