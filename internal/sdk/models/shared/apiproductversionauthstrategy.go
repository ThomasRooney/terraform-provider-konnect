// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type APIProductVersionAuthStrategy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (o *APIProductVersionAuthStrategy) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *APIProductVersionAuthStrategy) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
