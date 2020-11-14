// Code generated by goa v3.2.5, DO NOT EDIT.
//
// book HTTP server types
//
// Command:
// $ goa gen book/design

package server

import (
	book "book/gen/book"
	bookviews "book/gen/book/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "book" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// ID of the book
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of book
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the book
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Price of the book
	Price *uint32 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// UpdateRequestBody is the type of the "book" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Name of book
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the book
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Price of the book
	Price *uint32 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// CreateResponseBody is the type of the "book" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// ID of the book
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Name of book
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the book
	Description string `form:"description" json:"description" xml:"description"`
	// Price of the book
	Price uint32 `form:"price" json:"price" xml:"price"`
}

// ListResponseBody is the type of the "book" service "list" endpoint HTTP
// response body.
type ListResponseBody []*BookResponse

// RemoveNotFoundResponseBody is the type of the "book" service "remove"
// endpoint HTTP response body for the "not-found" error.
type RemoveNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// BookResponse is used to define fields on response body types.
type BookResponse struct {
	// ID of the book
	ID uint32 `form:"id" json:"id" xml:"id"`
	// Name of book
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the book
	Description string `form:"description" json:"description" xml:"description"`
	// Price of the book
	Price uint32 `form:"price" json:"price" xml:"price"`
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "book" service.
func NewCreateResponseBody(res *bookviews.BookView) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:          *res.ID,
		Name:        *res.Name,
		Description: *res.Description,
		Price:       *res.Price,
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "book" service.
func NewListResponseBody(res []*book.Book) ListResponseBody {
	body := make([]*BookResponse, len(res))
	for i, val := range res {
		body[i] = marshalBookBookToBookResponse(val)
	}
	return body
}

// NewRemoveNotFoundResponseBody builds the HTTP response body from the result
// of the "remove" endpoint of the "book" service.
func NewRemoveNotFoundResponseBody(res *goa.ServiceError) *RemoveNotFoundResponseBody {
	body := &RemoveNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateBook builds a book service create endpoint payload.
func NewCreateBook(body *CreateRequestBody) *book.Book {
	v := &book.Book{
		ID:          *body.ID,
		Name:        *body.Name,
		Description: *body.Description,
		Price:       *body.Price,
	}

	return v
}

// NewUpdateBook builds a book service update endpoint payload.
func NewUpdateBook(body *UpdateRequestBody, id uint32) *book.Book {
	v := &book.Book{
		Name:        *body.Name,
		Description: *body.Description,
		Price:       *body.Price,
	}
	v.ID = id

	return v
}

// NewRemovePayload builds a book service remove endpoint payload.
func NewRemovePayload(id uint32) *book.RemovePayload {
	v := &book.RemovePayload{}
	v.ID = id

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 100, false))
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 100, false))
		}
	}
	return
}
