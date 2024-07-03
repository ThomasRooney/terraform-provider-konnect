// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateSystemAccount struct {
	// Description of the system account. Useful when the system account name is not sufficient to differentiate one system account from another.
	Description string `json:"description"`
	// The system account is managed by Konnect (true/false).
	KonnectManaged *bool `json:"konnect_managed,omitempty"`
	// Name of the system account.
	Name string `json:"name"`
}

func (o *CreateSystemAccount) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *CreateSystemAccount) GetKonnectManaged() *bool {
	if o == nil {
		return nil
	}
	return o.KonnectManaged
}

func (o *CreateSystemAccount) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}
