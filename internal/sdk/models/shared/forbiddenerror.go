// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ForbiddenError - standard error
type ForbiddenError struct {
	Status   any `json:"status"`
	Title    any `json:"title"`
	Type     any `json:"type,omitempty"`
	Instance any `json:"instance"`
	Detail   any `json:"detail"`
}

func (o *ForbiddenError) GetStatus() any {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ForbiddenError) GetTitle() any {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ForbiddenError) GetType() any {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ForbiddenError) GetInstance() any {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *ForbiddenError) GetDetail() any {
	if o == nil {
		return nil
	}
	return o.Detail
}
