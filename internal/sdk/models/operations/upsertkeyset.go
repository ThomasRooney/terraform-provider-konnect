// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpsertKeySetRequest struct {
	// ID of the KeySet to lookup
	KeySetID string `pathParam:"style=simple,explode=false,name=KeySetId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Description of the KeySet
	KeySet shared.KeySetInput `request:"mediaType=application/json"`
}

func (o *UpsertKeySetRequest) GetKeySetID() string {
	if o == nil {
		return ""
	}
	return o.KeySetID
}

func (o *UpsertKeySetRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpsertKeySetRequest) GetKeySet() shared.KeySetInput {
	if o == nil {
		return shared.KeySetInput{}
	}
	return o.KeySet
}

type UpsertKeySetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted KeySet
	KeySet *shared.KeySet
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpsertKeySetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertKeySetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertKeySetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertKeySetResponse) GetKeySet() *shared.KeySet {
	if o == nil {
		return nil
	}
	return o.KeySet
}

func (o *UpsertKeySetResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
