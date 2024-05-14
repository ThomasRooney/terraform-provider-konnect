// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type GetTransitGatewayRequest struct {
	// The network to operate on.
	NetworkID string `pathParam:"style=simple,explode=false,name=networkId"`
	// The ID of the transit gateway to operate on.
	TransitGatewayID string `pathParam:"style=simple,explode=false,name=transitGatewayId"`
}

func (o *GetTransitGatewayRequest) GetNetworkID() string {
	if o == nil {
		return ""
	}
	return o.NetworkID
}

func (o *GetTransitGatewayRequest) GetTransitGatewayID() string {
	if o == nil {
		return ""
	}
	return o.TransitGatewayID
}

type GetTransitGatewayResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Response format for retrieving a transit gateway.
	TransitGateway *shared.TransitGateway
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetTransitGatewayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransitGatewayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransitGatewayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransitGatewayResponse) GetTransitGateway() *shared.TransitGateway {
	if o == nil {
		return nil
	}
	return o.TransitGateway
}

func (o *GetTransitGatewayResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *GetTransitGatewayResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}

func (o *GetTransitGatewayResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
