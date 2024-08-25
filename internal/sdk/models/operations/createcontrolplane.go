// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateControlPlaneResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response to creating a control plane.
	ControlPlane *shared.ControlPlane
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Unauthenticated
	UnauthorizedError *shared.UnauthorizedError
	// Permission denied
	ForbiddenError *shared.ForbiddenError
	// Conflict
	ConflictError *shared.ConflictError
	// Internal Server Error
	InternalServerError *shared.InternalServerError
	// Service Unavailable
	ServiceUnavailable *shared.ServiceUnavailable
}

func (o *CreateControlPlaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateControlPlaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateControlPlaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateControlPlaneResponse) GetControlPlane() *shared.ControlPlane {
	if o == nil {
		return nil
	}
	return o.ControlPlane
}

func (o *CreateControlPlaneResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *CreateControlPlaneResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateControlPlaneResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *CreateControlPlaneResponse) GetConflictError() *shared.ConflictError {
	if o == nil {
		return nil
	}
	return o.ConflictError
}

func (o *CreateControlPlaneResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateControlPlaneResponse) GetServiceUnavailable() *shared.ServiceUnavailable {
	if o == nil {
		return nil
	}
	return o.ServiceUnavailable
}
