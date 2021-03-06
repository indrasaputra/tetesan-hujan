package entity

// Bookmark defines the data type to save a bookmark.
type Bookmark struct {
	// CollectionName defines in which collection the bookmark is saved.
	CollectionName string `json:"collection_name"`
	// URL defines the url of the bookmark.
	URL string `json:"url"`
}

// Raindrop defines the data structure to save a bookmark in raindrop.io.
type Raindrop struct {
	// Title defines the raindrop's title.
	Title string `json:"title"`
	// Excerpt defines the raindrop's excerpt.
	Excerpt string `json:"excerpt"`
	// Link defines the bookmark's link.
	// Link is better to be set via parsedURL.meta.canonical.
	Link string `json:"link"`
	// CollectionID defines in which collection the raindrop will save the bookmark.
	CollectionID int64 `json:"collectionID"`
}

// ParsedURL defines the result of parsed URL from Raindrop.
type ParsedURL struct {
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
