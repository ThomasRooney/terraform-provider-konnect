// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ListProviderAccountsResponse struct {
	Data []ProviderAccount `json:"data"`
	// returns the pagination information
	Meta PaginatedMeta `json:"meta"`
}

func (o *ListProviderAccountsResponse) GetData() []ProviderAccount {
	if o == nil {
		return []ProviderAccount{}
	}
	return o.Data
}

func (o *ListProviderAccountsResponse) GetMeta() PaginatedMeta {
	if o == nil {
		return PaginatedMeta{}
	}
	return o.Meta
}
