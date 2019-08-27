// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListJokeFormatedOKCode is the HTTP code returned for type ListJokeFormatedOK
const ListJokeFormatedOKCode int = 200

/*ListJokeFormatedOK Returns requested Entities.

swagger:response listJokeFormatedOK
*/
type ListJokeFormatedOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListJokeFormatedOK creates ListJokeFormatedOK with default headers values
func NewListJokeFormatedOK() *ListJokeFormatedOK {

	return &ListJokeFormatedOK{}
}

// WithPayload adds the payload to the list joke formated o k response
func (o *ListJokeFormatedOK) WithPayload(payload interface{}) *ListJokeFormatedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list joke formated o k response
func (o *ListJokeFormatedOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListJokeFormatedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
