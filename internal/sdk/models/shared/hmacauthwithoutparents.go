// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type HMACAuthWithoutParents struct {
	Secret   *string  `json:"secret,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Username *string  `json:"username,omitempty"`
}

func (o *HMACAuthWithoutParents) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *HMACAuthWithoutParents) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *HMACAuthWithoutParents) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
