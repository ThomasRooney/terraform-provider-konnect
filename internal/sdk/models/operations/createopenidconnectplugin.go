// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateOpenidconnectPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID            string                            `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateOpenidConnectPlugin *shared.CreateOpenidConnectPlugin `request:"mediaType=application/json"`
}

func (o *CreateOpenidconnectPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateOpenidconnectPluginRequest) GetCreateOpenidConnectPlugin() *shared.CreateOpenidConnectPlugin {
	if o == nil {
		return nil
	}
	return o.CreateOpenidConnectPlugin
}

type CreateOpenidconnectPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created OpenidConnect plugin
	OpenidConnectPlugin *shared.OpenidConnectPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateOpenidconnectPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateOpenidconnectPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateOpenidconnectPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateOpenidconnectPluginResponse) GetOpenidConnectPlugin() *shared.OpenidConnectPlugin {
	if o == nil {
		return nil
	}
	return o.OpenidConnectPlugin
}

func (o *CreateOpenidconnectPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
