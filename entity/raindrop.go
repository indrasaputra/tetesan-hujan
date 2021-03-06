package entity

// Bookmark defines the data type to save a bookmark.
type Bookmark struct {
	// CollectionName defines in which collection the bookmark is saved.
	CollectionName string `json:"collection_name"`
	// URL defines the url of the bookmark.
	URL string `json:"url"`
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
