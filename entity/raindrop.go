package entity

// Raindrop defines the data type to save a bookmark.
// It is named Raindrop because it is integrated with raindrop.io
type Raindrop struct {
	// CollectionName defines in which collection the bookmark is saved.
	CollectionName string `json:"collection_name"`
	// Link defines the link of the bookmark.
	Link string `json:"link"`
}

// Collection defines the data tyope for a collection.
type Collection struct {
	// ID defines the collection's ID.
	ID int64 `json:"_id"`
	// Name defines collection's name.
	Name string `json:"name"`
}
