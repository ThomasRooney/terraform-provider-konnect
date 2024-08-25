// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

var CreateServerlessCloudGatewayServerList = []string{
	"https://global.api.konghq.com/",
}

type CreateServerlessCloudGatewayResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response to creating a serverless cloud gateway.
	ServerlessCloudGateway *shared.ServerlessCloudGateway
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
	// Forbidden
	ForbiddenError *shared.ForbiddenError
}

func (o *CreateServerlessCloudGatewayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateServerlessCloudGatewayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateServerlessCloudGatewayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateServerlessCloudGatewayResponse) GetServerlessCloudGateway() *shared.ServerlessCloudGateway {
	if o == nil {
		return nil
	}
	return o.ServerlessCloudGateway
}

func (o *CreateServerlessCloudGatewayResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *CreateServerlessCloudGatewayResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}

func (o *CreateServerlessCloudGatewayResponse) GetForbiddenError() *shared.ForbiddenError {
	if o == nil {
		return nil
	}
	return o.ForbiddenError
}
