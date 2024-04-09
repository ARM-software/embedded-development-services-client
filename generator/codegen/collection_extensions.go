package codegen

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/ARM-software/golang-utils/utils/commonerrors"
)

const (
	schemaPrefix   = "#/components/schemas/"
	jsonMIME       = "application/json"
	collectionFlag = "x-collection"
	redactFlag     = "x-redact"
	// Messages are a special case as they are a feed rather than a normal collection
	notificationFeedRef = "NotificationFeed"
)

var ignoreItems = []string{
	// Deprecation notices don't have a collection
	"#/components/schemas/EndpointDeprecationNotice",
	// VendorItems don't follow the pattern in the schema
	"#/components/schemas/VendorItem",
	// The following don't have corresponding clients or haven't been updated
	"#/components/schemas/GenericWorkerItem",
	"#/components/schemas/GenericWorkJobItem",
	"#/components/schemas/PATItem",
	"#/components/schemas/PATCollection",
}

var ignoreCollections = []string{
	// SimpleCollection is an underlying type
	"#/components/schemas/SimpleCollection",
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
		err = fmt.Errorf("%w: schemaVal must be defined", commonerrors.ErrUndefined)
		return
	}

	embeddedRef := schemaVal.Properties["_embedded"]
	if embeddedRef == nil {
		err = fmt.Errorf("%w: The schema for '%s' does not contain embedded objects so their type cannot be determined", commonerrors.ErrNotFound, collectionRef)
		return
	}

	embeddedSchemaVal := embeddedRef.Value
	if embeddedSchemaVal == nil {
		err = fmt.Errorf("%w: The schema response for collection references '%s' exists but has no value", commonerrors.ErrUndefined, collectionRef)
		return
	}

	underlyingItem := embeddedSchemaVal.Properties["item"]
	if underlyingItem == nil {
		err = fmt.Errorf("%w: The embedded objects for schema for '%s' is missing the item field in its embedded item properties", commonerrors.ErrNotFound, collectionRef)
		return
	}

	underlyingItemValue := underlyingItem.Value
	if underlyingItemValue == nil {
		err = fmt.Errorf("%w: The embedded objects for schema for '%s' contains the item field but it is undefined", commonerrors.ErrUndefined, collectionRef)
		return
	}

	if underlyingItemValue.Items == nil {
		err = fmt.Errorf("%w: The embedded objects for schema for '%s' is missing the items field that was expected in a collection field", commonerrors.ErrInvalid, collectionRef)
		return
	}

	itemRef = underlyingItem.Value.Items.Ref
	return
}

func getResponseFromPathInfo(pathInfo *openapi3.PathItem, endpoint string, status int) (resp *openapi3.Response, err error) {
	if pathInfo == nil {
		err = fmt.Errorf("%w: pathInfo must be defined", commonerrors.ErrUndefined)
		return
	}

	response200 := pathInfo.Get.Responses.Get(status)
	if response200 == nil {
		err = fmt.Errorf("%w: The %v response code is missing from object with endpoint '%s'", commonerrors.ErrNotFound, status, endpoint)
		return
	}

	resp = response200.Value
	if resp == nil {
		err = fmt.Errorf("%w: The %v response code exists in object with endpoint '%s' but is missing a value", commonerrors.ErrUndefined, status, endpoint)
		return
	}

	return
}

// message collections don't use the same structure as other collections
func parseMessageCollection(schemaVal *openapi3.Schema, endpoint string) (messagesItemRef string, err error) {
	if schemaVal == nil {
		err = fmt.Errorf("%w: schemaVal must be defined", commonerrors.ErrUndefined)
		return
	}

	messagesRef := schemaVal.Properties["messages"]
	if messagesRef == nil {
		err = fmt.Errorf("%w: The the message endpoint '%v' does not contain a messages field ", commonerrors.ErrNotFound, endpoint)
		return
	}

	messagesValue := messagesRef.Value
	if messagesValue == nil {
		err = fmt.Errorf("%w: The messages object for '%v' contains the messages field but it is undefined", commonerrors.ErrUndefined, endpoint)
		return
	}

	if messagesValue.Items == nil {
		err = fmt.Errorf("%w: The messages object for '%v' is missing the items field that was expected in a collection field", commonerrors.ErrInvalid, endpoint)
		return
	}

	messagesItemRef = trimRefPrefix(messagesValue.Items.Ref)
	return
}

func getCollectionSchema(swagger *openapi3.T, appJSON *openapi3.MediaType, endpoint string) (schemaVal *openapi3.Schema, collectionRef string, err error) {
	if swagger == nil {
		err = fmt.Errorf("%w: swagger must be defined", commonerrors.ErrUndefined)
		return
	}
	if appJSON == nil {
		err = fmt.Errorf("%w: appJSON must be defined", commonerrors.ErrUndefined)
		return
	}

	schema := appJSON.Schema
	if schema == nil {
		err = fmt.Errorf("%w: The object with endpoint '%s' does not have a reference to the schema denoting the type of the response", commonerrors.ErrNotFound, endpoint)
		return
	}
	ref := schema.Ref
	if ref == "" {
		err = fmt.Errorf("%w: The object with endpoint '%s' has a field that should contain a reference the schema denoting the type of the response, but it is empty", commonerrors.ErrEmpty, endpoint)
		return
	}

	collectionRef = trimRefPrefix(ref)
	schema, found := swagger.Components.Schemas[collectionRef]
	if !found {
		err = fmt.Errorf("%w: The reference at '%s' does not contain the type of the endpoint object (%v)", commonerrors.ErrNotFound, endpoint, collectionRef)
		return
	}

	schemaVal = schema.Value
	if schemaVal == nil {
		err = fmt.Errorf("%w: The schema response for collection references '%s' exists but has no value", commonerrors.ErrUndefined, collectionRef)
		return
	}

	return
}

func newCollection(collectionRef, itemRef string) Collection {
	return Collection{
		CollectionRef: collectionRef,
		ItemRef:       trimRefPrefix(itemRef),
		ModelRef:      fmt.Sprintf("%sModel", strings.TrimSuffix(collectionRef, "Collection")),
		IteratorRef:   fmt.Sprintf("%sIterator", strings.TrimSuffix(collectionRef, "Collection")),
	}
}

func newMessageCollection(itemRef string) MessageCollection {
	return MessageCollection{
		ItemRef:     trimRefPrefix(itemRef),
		IteratorRef: fmt.Sprintf("%sIterator", strings.TrimSuffix(itemRef, "Object")),
	}
}
func newNotificationFeedCollection(itemRef string) NotificationFeedCollection {
	return NotificationFeedCollection{
		ItemRef: trimRefPrefix(itemRef),
	}
}

func GetCollections(swagger *openapi3.T) (collections CollectionParams, err error) {
	collectionSet := mapset.NewSet[Collection]()
	messageCollectionsSet := mapset.NewSet[MessageCollection]()
	notificationFeedCollectionsSet := mapset.NewSet[NotificationFeedCollection]()

	for endpoint, pathInfo := range swagger.Paths {
		var isCollection bool

		// if no get field then not a collection
		if pathInfo.Get != nil {
			response200Val, subErr := getResponseFromPathInfo(pathInfo, endpoint, 200)
			if subErr != nil {
				err = subErr
				return
			}

			var isRedacted bool
			if c, ok := pathInfo.Get.ExtensionProps.Extensions[redactFlag].(json.RawMessage); ok {
				err = json.Unmarshal(c, &isRedacted)
				if err != nil {
					return
				}
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

				if c, ok := schemaVal.ExtensionProps.Extensions[collectionFlag].(json.RawMessage); ok {
					err = json.Unmarshal(c, &isCollection)
					if err != nil {
						return
					}

					if !isCollection || shouldIgnoreCollection(collectionRef) {
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

					collectionSet.Add(newCollection(collectionRef, itemRef))
				}

				if isMessagesCollection := strings.HasSuffix(endpoint, "messages"); isMessagesCollection {
					if strings.Contains(collectionRef, notificationFeedRef) {
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

	collections = CollectionParams{
		collectionSlice,
		messageCollectionSlice,
		notificationFeedCollectionSlice,
	}

	return
}
