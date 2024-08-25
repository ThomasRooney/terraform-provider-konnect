// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdatePrefunctionPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID          string                          `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreatePreFunctionPlugin *shared.CreatePreFunctionPlugin `request:"mediaType=application/json"`
}

func (o *UpdatePrefunctionPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdatePrefunctionPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdatePrefunctionPluginRequest) GetCreatePreFunctionPlugin() *shared.CreatePreFunctionPlugin {
	if o == nil {
		return nil
	}
	return o.CreatePreFunctionPlugin
}

type UpdatePrefunctionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// PreFunction plugin
	PreFunctionPlugin *shared.PreFunctionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdatePrefunctionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePrefunctionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePrefunctionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePrefunctionPluginResponse) GetPreFunctionPlugin() *shared.PreFunctionPlugin {
	if o == nil {
		return nil
	}
	return o.PreFunctionPlugin
}

func (o *UpdatePrefunctionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
