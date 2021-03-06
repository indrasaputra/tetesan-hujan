package entity

// Raindrop defines the data type to save a bookmark.
// It is named Raindrop because it is integrated with raindrop.io
type Raindrop struct {
	// CollectionName defines in which collection the bookmark is saved.
	CollectionName string `json:"collection_name"`
	// Link defines the link of the bookmark.
	Link string `json:"link"`
}

// ParsedURL defines the result of parsed URL from Raindrop.
type ParsedURL struct {
	// Result defines the parse URL result.
	Result bool `json:"result"`
	// Error defines the parse URL error.
	Error string `json:"error"`
	Item  struct {
		// Title defines the title of the URL.
		Title string `json:"title"`
		// Excerpt defines the excerpt of the URL.
		Excerpt string `json:"excerpt"`
		// Meta defines the metadata of the item
		Meta struct {
			Canonical string `json:"canonical"`
		} `json:"meta"`
	} `json:"item"`
}

// Collection defines the data tyope for a collection.
type Collection struct {
	// ID defines the collection's ID.
	ID int64 `json:"_id"`
	// Name defines collection's name.
	Name string `json:"name"`
}
