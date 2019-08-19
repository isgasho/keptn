// Code generated by go-swagger; DO NOT EDIT.

package stage_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/configuration-service/models"
)

// DeleteProjectProjectNameStageStageNameResourceResourceURINoContentCode is the HTTP code returned for type DeleteProjectProjectNameStageStageNameResourceResourceURINoContent
const DeleteProjectProjectNameStageStageNameResourceResourceURINoContentCode int = 204

/*DeleteProjectProjectNameStageStageNameResourceResourceURINoContent Success. Stage resource has been deleted. Response does not have a body.

swagger:response deleteProjectProjectNameStageStageNameResourceResourceUriNoContent
*/
type DeleteProjectProjectNameStageStageNameResourceResourceURINoContent struct {
}

// NewDeleteProjectProjectNameStageStageNameResourceResourceURINoContent creates DeleteProjectProjectNameStageStageNameResourceResourceURINoContent with default headers values
func NewDeleteProjectProjectNameStageStageNameResourceResourceURINoContent() *DeleteProjectProjectNameStageStageNameResourceResourceURINoContent {

	return &DeleteProjectProjectNameStageStageNameResourceResourceURINoContent{}
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURINoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequestCode is the HTTP code returned for type DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest
const DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequestCode int = 400

/*DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest Failed. Stage resource could not be deleted.

swagger:response deleteProjectProjectNameStageStageNameResourceResourceUriBadRequest
*/
type DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest creates DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest with default headers values
func NewDeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest() *DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest {

	return &DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest{}
}

// WithPayload adds the payload to the delete project project name stage stage name resource resource Uri bad request response
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest) WithPayload(payload *models.Error) *DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project project name stage stage name resource resource Uri bad request response
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURIBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteProjectProjectNameStageStageNameResourceResourceURIDefault Error

swagger:response deleteProjectProjectNameStageStageNameResourceResourceUriDefault
*/
type DeleteProjectProjectNameStageStageNameResourceResourceURIDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteProjectProjectNameStageStageNameResourceResourceURIDefault creates DeleteProjectProjectNameStageStageNameResourceResourceURIDefault with default headers values
func NewDeleteProjectProjectNameStageStageNameResourceResourceURIDefault(code int) *DeleteProjectProjectNameStageStageNameResourceResourceURIDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteProjectProjectNameStageStageNameResourceResourceURIDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete project project name stage stage name resource resource URI default response
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURIDefault) WithStatusCode(code int) *DeleteProjectProjectNameStageStageNameResourceResourceURIDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete project project name stage stage name resource resource URI default response
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURIDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete project project name stage stage name resource resource URI default response
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURIDefault) WithPayload(payload *models.Error) *DeleteProjectProjectNameStageStageNameResourceResourceURIDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete project project name stage stage name resource resource URI default response
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURIDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteProjectProjectNameStageStageNameResourceResourceURIDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}