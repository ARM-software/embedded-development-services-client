package codegen

type Collection struct {
	CollectionRef string
	ItemRef       string
	ModelRef      string
	IteratorRef   string
}

type Collections = []Collection

type MessageCollection struct {
	ItemRef     string
	IteratorRef string
}

type MessageCollections []MessageCollection

type NotificationFeedCollection struct {
	ItemRef string
}

type NotificationFeedCollections []NotificationFeedCollection

type CollectionParams = struct {
	Collections
	MessageCollections
	NotificationFeedCollections
}
