// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateIprestrictionPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID            string                            `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateIPRestrictionPlugin *shared.CreateIPRestrictionPlugin `request:"mediaType=application/json"`
}

func (o *CreateIprestrictionPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateIprestrictionPluginRequest) GetCreateIPRestrictionPlugin() *shared.CreateIPRestrictionPlugin {
	if o == nil {
		return nil
	}
	return o.CreateIPRestrictionPlugin
}

type CreateIprestrictionPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created IpRestriction plugin
	IPRestrictionPlugin *shared.IPRestrictionPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateIprestrictionPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateIprestrictionPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateIprestrictionPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateIprestrictionPluginResponse) GetIPRestrictionPlugin() *shared.IPRestrictionPlugin {
	if o == nil {
		return nil
	}
	return o.IPRestrictionPlugin
}

func (o *CreateIprestrictionPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
