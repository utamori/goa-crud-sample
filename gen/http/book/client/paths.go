// Code generated by goa v3.2.5, DO NOT EDIT.
//
// HTTP request path constructors for the book service.
//
// Command:
// $ goa gen book/design

package client

import (
	"fmt"
)

// CreateBookPath returns the URL path to the book service create HTTP endpoint.
func CreateBookPath() string {
	return "/"
}

// ListBookPath returns the URL path to the book service list HTTP endpoint.
func ListBookPath() string {
	return "/books"
}

// UpdateBookPath returns the URL path to the book service update HTTP endpoint.
func UpdateBookPath(id uint32) string {
	return fmt.Sprintf("/book/%v", id)
}

// RemoveBookPath returns the URL path to the book service remove HTTP endpoint.
func RemoveBookPath(id uint32) string {
	return fmt.Sprintf("/book/%v", id)
}
