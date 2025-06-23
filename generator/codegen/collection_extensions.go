package codegen

import (
	"fmt"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
)

type Collection struct {
	CollectionRef string
	ItemRef       string
	ModelRef      string
	IteratorRef   string
	NoPagination  bool
}

type Collections = []Collection

type MessageCollection struct {
	ItemRef     string
	IteratorRef string
}

type MessageCollections []MessageCollection

type NotificationFeedCollection struct {
	ItemRef     string
	IteratorRef string
}

type NotificationFeedCollections []NotificationFeedCollection

type CollectionParams = struct {
	Collections
	JobItems
	MessageCollections
	NotificationFeedCollections
}

type collectionOptions struct {
	noPagination bool
}

type CollectionOption interface {
	apply(*collectionOptions)
}

type noPaginationOption bool

func (n noPaginationOption) apply(opts *collectionOptions) {
	opts.noPagination = bool(n)
}

func WithNoPagination(n bool) CollectionOption {
	return noPaginationOption(n)
}

const (
	// Messages are a special case as they are a feed rather than a normal collection
	notificationFeedRef = "NotificationFeed"
	messageItemRef      = "MessageItem"
)

var ignoreItems = []string{
	// Deprecation notices don't have a collection
	"#/components/schemas/EndpointDeprecationNotice",
	// VendorItems don't follow the pattern in the schema
	"#/components/schemas/VendorItem",
	// The following don't have corresponding clients or haven't been updated
	"#/components/schemas/PATItem",
	"#/components/schemas/PATCollection",
}

var ignoreCollections = []string{
	// SimpleCollection is an underlying type
	"#/components/schemas/SimpleCollection",
}

func AddCollectionsToParams(d *Data) (err error) {
	return AddValuesToParams(d, func(swagger *openapi3.T) (interface{}, error) { return GetCollections(swagger) }, "entities.go")
}

func trimRefPrefix(ref string) string {
	return strings.TrimPrefix(ref, schemaPrefix)
}

func shouldIgnoreItem(item string) bool {
	return slices.Contains(ignoreItems, item) ||
		slices.Contains(ignoreItems, trimRefPrefix(item)) ||
		slices.Contains(ignoreItems, fmt.Sprintf("%s%s", schemaPrefix, item))
}

func shouldIgnoreCollection(collection string) bool {
	return slices.Contains(ignoreCollections, collection) ||
		slices.Contains(ignoreCollections, trimRefPrefix(collection)) ||
		slices.Contains(ignoreCollections, fmt.Sprintf("%s%s", schemaPrefix, collection))
}

func extractItemRef(schemaVal *openapi3.Schema, collectionRef string) (itemRef string, err error) {
	if schemaVal == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "schemaVal must be defined")
		return
	}

	embeddedRef := schemaVal.Properties["_embedded"]
	if embeddedRef == nil {
		err = commonerrors.Newf(commonerrors.ErrNotFound, "the schema for '%s' does not contain embedded objects so their type cannot be determined", collectionRef)
		return
	}

	embeddedSchemaVal := embeddedRef.Value
	if embeddedSchemaVal == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "the schema response for collection references '%s' exists but has no value", collectionRef)
		return
	}

	underlyingItem := embeddedSchemaVal.Properties["item"]
	if underlyingItem == nil {
		err = commonerrors.Newf(commonerrors.ErrNotFound, "the embedded objects for schema for '%s' is missing the item field in its embedded item properties", collectionRef)
		return
	}

	underlyingItemValue := underlyingItem.Value
	if underlyingItemValue == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "the embedded objects for schema for '%s' contains the item field but it is undefined", collectionRef)
		return
	}

	if underlyingItemValue.Items == nil {
		err = commonerrors.Newf(commonerrors.ErrInvalid, "the embedded objects for schema for '%s' is missing the items field that was expected in a collection field", collectionRef)
		return
	}

	itemRef = underlyingItem.Value.Items.Ref
	return
}

func getResponseFromPathInfo(pathInfo *openapi3.PathItem, endpoint string, status int) (resp *openapi3.Response, err error) {
	if pathInfo == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "pathInfo must be defined")
		return
	}

	response200 := pathInfo.Get.Responses.Get(status)
	if response200 == nil {
		err = commonerrors.Newf(commonerrors.ErrNotFound, "the %v response code is missing from object with endpoint '%s'", status, endpoint)
		return
	}

	resp = response200.Value
	if resp == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "the %v response code exists in object with endpoint '%s' but is missing a value", status, endpoint)
		return
	}

	return
}

// message collections don't use the same structure as other collections
func parseMessageCollection(schemaVal *openapi3.Schema, endpoint string) (messagesItemRef string, err error) {
	if schemaVal == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "schemaVal must be defined")
		return
	}

	messagesRef := schemaVal.Properties["messages"]
	if messagesRef == nil {
		err = commonerrors.Newf(commonerrors.ErrNotFound, "the the message endpoint '%v' does not contain a messages field ", endpoint)
		return
	}

	messagesValue := messagesRef.Value
	if messagesValue == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "the messages object for '%v' contains the messages field but it is undefined", endpoint)
		return
	}

	if messagesValue.Items == nil {
		err = commonerrors.Newf(commonerrors.ErrInvalid, "the messages object for '%v' is missing the items field that was expected in a collection field", endpoint)
		return
	}

	messagesItemRef = trimRefPrefix(messagesValue.Items.Ref)
	return
}

func getCollectionSchema(swagger *openapi3.T, appJSON *openapi3.MediaType, endpoint string) (schemaVal *openapi3.Schema, collectionRef string, err error) {
	if swagger == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "swagger must be defined")
		return
	}
	if appJSON == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "appJSON must be defined")
		return
	}

	schema := appJSON.Schema
	if schema == nil {
		err = commonerrors.Newf(commonerrors.ErrNotFound, "the object with endpoint '%s' does not have a reference to the schema denoting the type of the response", endpoint)
		return
	}
	ref := schema.Ref
	if ref == "" {
		err = commonerrors.Newf(commonerrors.ErrEmpty, "the object with endpoint '%s' has a field that should contain a reference the schema denoting the type of the response, but it is empty", endpoint)
		return
	}

	collectionRef = trimRefPrefix(ref)
	schema, found := swagger.Components.Schemas[collectionRef]
	if !found {
		err = commonerrors.Newf(commonerrors.ErrNotFound, "the reference at '%s' does not contain the type of the endpoint object (%v)", endpoint, collectionRef)
		return
	}

	schemaVal = schema.Value
	if schemaVal == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "the schema response for collection references '%s' exists but has no value", collectionRef)
		return
	}

	return
}

func newCollection(collectionRef, itemRef string, opts ...CollectionOption) Collection {
	options := collectionOptions{
		noPagination: false,
	}

	for _, opt := range opts {
		opt.apply(&options)
	}

	return Collection{
		CollectionRef: collectionRef,
		ItemRef:       trimRefPrefix(itemRef),
		ModelRef:      fmt.Sprintf("%sModel", strings.TrimSuffix(collectionRef, "Collection")),
		IteratorRef:   fmt.Sprintf("%sIterator", strings.TrimSuffix(collectionRef, "Collection")),
		NoPagination:  options.noPagination,
	}
}

func newMessageCollection(itemRef string) MessageCollection {
	return MessageCollection{
		ItemRef:     trimRefPrefix(itemRef),
		IteratorRef: fmt.Sprintf("%sIterator", strings.TrimSuffix(itemRef, "Object")),
	}
}

func newNotificationFeedCollection(itemRef string) NotificationFeedCollection {
	var iteratorRef string
	switch {
	case strings.Contains(itemRef, notificationFeedRef):
		iteratorRef = "NewNotificationMessageIterator"
	case strings.Contains(itemRef, messageItemRef):
		iteratorRef = "NewMessageIterator"
	}
	return NotificationFeedCollection{
		ItemRef:     trimRefPrefix(itemRef),
		IteratorRef: iteratorRef,
	}
}

func getStandaloneJobItems(swagger *openapi3.T, collections []Collection) (standalone JobItems, err error) {
	jobsParams, subErr := GetJobItems(swagger)
	if subErr != nil {
		err = subErr
		return
	}

	jobSet := mapset.NewSet[JobItem]()
	for _, jobItem := range jobsParams.JobItems {
		jobSet.Add(jobItem)
	}

	for _, collection := range collections {
		jobTarget := JobItem{collection.ItemRef}
		if jobSet.Contains(jobTarget) {
			jobSet.Remove(jobTarget)
		}
	}

	standalone = jobSet.ToSlice()

	return
}

func GetCollections(swagger *openapi3.T) (collections CollectionParams, err error) {
	collectionSet := mapset.NewSet[Collection]()
	messageCollectionsSet := mapset.NewSet[MessageCollection]()
	notificationFeedCollectionsSet := mapset.NewSet[NotificationFeedCollection]()

	for endpoint, pathInfo := range swagger.Paths {
		// if no get field then not a collection
		if pathInfo.Get != nil {
			response200Val, subErr := getResponseFromPathInfo(pathInfo, endpoint, 200)
			if subErr != nil {
				err = subErr
				return
			}

			isRedacted, subErr := isExtensionFlagSet(pathInfo.Get.ExtensionProps, redactFlag)
			if subErr != nil {
				err = subErr
				return
			}
			if isRedacted {
				continue
			}

			// if get not "application/json" then it is "application/octet-stream" and collections don't have that
			if appJSON := response200Val.Content.Get(jsonMIME); appJSON != nil {
				schemaVal, collectionRef, subErr := getCollectionSchema(swagger, appJSON, endpoint)
				if subErr != nil {
					err = subErr
					return
				}

				isCollection, subErr := isExtensionFlagSet(schemaVal.ExtensionProps, collectionFlag)
				if subErr != nil {
					err = subErr
					return
				}

				if isCollection {
					if shouldIgnoreCollection(collectionRef) {
						continue
					}

					itemRef, subErr := extractItemRef(schemaVal, collectionRef)
					if subErr != nil {
						err = subErr
						return
					}

					if shouldIgnoreItem(itemRef) {
						continue
					}

					isNoPaginationCollection, subErr := isExtensionFlagSet(schemaVal.ExtensionProps, noPaginationFlag)
					if subErr != nil {
						err = subErr
						return
					}

					if isNoPaginationCollection {
						collectionSet.Add(newCollection(collectionRef, itemRef, WithNoPagination(true)))
					} else {
						collectionSet.Add(newCollection(collectionRef, itemRef))
					}
				}

				if isMessagesCollection := strings.HasSuffix(endpoint, "messages"); isMessagesCollection {
					if strings.Contains(collectionRef, notificationFeedRef) || strings.Contains(collectionRef, messageItemRef) {
						if shouldIgnoreItem(collectionRef) {
							continue
						}

						notificationFeedCollectionsSet.Add(newNotificationFeedCollection(collectionRef))
					}

					messagesRef, subErr := parseMessageCollection(schemaVal, collectionRef)
					if subErr != nil {
						err = subErr
						return
					}

					if shouldIgnoreItem(messagesRef) {
						continue
					}

					messageCollectionsSet.Add(newMessageCollection(messagesRef))
				}
			}
		}
	}

	// Sort slices to ensure order is the same between regeneration of file
	collectionSlice := collectionSet.ToSlice()
	messageCollectionSlice := messageCollectionsSet.ToSlice()
	notificationFeedCollectionSlice := notificationFeedCollectionsSet.ToSlice()
	slices.SortFunc(collectionSlice, func(a, b Collection) int { return strings.Compare(a.ItemRef, b.ItemRef) })
	slices.SortFunc(messageCollectionSlice, func(a, b MessageCollection) int { return strings.Compare(a.ItemRef, b.ItemRef) })
	slices.SortFunc(notificationFeedCollectionSlice, func(a, b NotificationFeedCollection) int { return strings.Compare(a.ItemRef, b.ItemRef) })

	standaloneJobItems, subErr := getStandaloneJobItems(swagger, collectionSlice)
	if subErr != nil {
		err = subErr
		return
	}

	collections = CollectionParams{
		collectionSlice,
		standaloneJobItems,
		messageCollectionSlice,
		notificationFeedCollectionSlice,
	}

	return
}
