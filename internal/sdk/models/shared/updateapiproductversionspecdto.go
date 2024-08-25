// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UpdateAPIProductVersionSpecDTO - The request schema for updating a specification for a version of an API product.
type UpdateAPIProductVersionSpecDTO struct {
	// The name of the API product version specification
	Name *string `json:"name,omitempty"`
	// The base64 encoded contents of the API product version specification
	Content *string `json:"content,omitempty"`
}

func (o *UpdateAPIProductVersionSpecDTO) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateAPIProductVersionSpecDTO) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}
