// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PortalOIDCConfig - Configuration properties for an OpenID Connect Identity Provider.
type PortalOIDCConfig struct {
	// Mappings from a portal developer atribute to an Identity Provider claim.
	ClaimMappings *PortalClaimMappings `json:"claim_mappings,omitempty"`
	ClientID      string               `json:"client_id"`
	Issuer        string               `json:"issuer"`
	Scopes        []string             `json:"scopes,omitempty"`
}

func (o *PortalOIDCConfig) GetClaimMappings() *PortalClaimMappings {
	if o == nil {
		return nil
	}
	return o.ClaimMappings
}

func (o *PortalOIDCConfig) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *PortalOIDCConfig) GetIssuer() string {
	if o == nil {
		return ""
	}
	return o.Issuer
}

func (o *PortalOIDCConfig) GetScopes() []string {
	if o == nil {
		return nil
	}
	return o.Scopes
}
