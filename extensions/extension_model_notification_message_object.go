// Package client defines an HTTP client for communicating with the web services.
// It includes the definition of request/response types as well as provides helpers for calling specific helpers.
package client

import (
	"errors"
)

// *************************************************************************************
// NOTE: this file is not generated.
// It extends the generated models.
// *************************************************************************************

// FetchType returns the resource type
func (o *NotificationMessageObject) FetchType() string {
	return "Message"
}

// NotificationMessageIterator defines an iterator over a message collection.
type NotificationMessageIterator struct {
	elements     []NotificationMessageObject
	currentIndex int
}

func (m *NotificationMessageIterator) HasNext() bool {
	return m.currentIndex < len(m.elements)
}

func (m *NotificationMessageIterator) GetNext() (item *interface{}, err error) {
	if m.currentIndex < 0 {
		err = errors.New("incorrect element index")
		return
	}
	if !m.HasNext() {
		err = errors.New("no more items")
		return
	}
	element := interface{}(m.elements[m.currentIndex])
	item = &element
	m.currentIndex++
	return
}

func NewNotificationMessagesIterator(elements []NotificationMessageObject) (IIterator, error) {
	return &NotificationMessageIterator{
		elements:     elements,
		currentIndex: 0,
	}, nil
}
