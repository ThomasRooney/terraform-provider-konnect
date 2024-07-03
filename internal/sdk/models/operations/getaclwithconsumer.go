// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type GetACLWithConsumerRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Consumer ID for nested entities
	ConsumerID string `pathParam:"style=simple,explode=false,name=ConsumerIdForNestedEntities"`
	// ID of the ACL to lookup
	ACLID string `pathParam:"style=simple,explode=false,name=ACLId"`
}

func (o *GetACLWithConsumerRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *GetACLWithConsumerRequest) GetConsumerID() string {
	if o == nil {
		return ""
	}
	return o.ConsumerID
}

func (o *GetACLWithConsumerRequest) GetACLID() string {
	if o == nil {
		return ""
	}
	return o.ACLID
}

type GetACLWithConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched ACL
	ACL *shared.ACL
}

func (o *GetACLWithConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetACLWithConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetACLWithConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetACLWithConsumerResponse) GetACL() *shared.ACL {
	if o == nil {
		return nil
	}
	return o.ACL
}
