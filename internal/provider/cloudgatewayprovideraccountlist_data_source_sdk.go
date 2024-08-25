// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *CloudGatewayProviderAccountListDataSourceModel) RefreshFromSharedListProviderAccountsResponse(resp *shared.ListProviderAccountsResponse) {
	if resp != nil {
		r.Data = []tfTypes.ProviderAccount{}
		if len(r.Data) > len(resp.Data) {
			r.Data = r.Data[:len(resp.Data)]
		}
		for dataCount, dataItem := range resp.Data {
			var data1 tfTypes.ProviderAccount
			data1.CreatedAt = types.StringValue(dataItem.CreatedAt.Format(time.RFC3339Nano))
			data1.ID = types.StringValue(dataItem.ID)
			data1.Provider = types.StringValue(string(dataItem.Provider))
			data1.ProviderAccountID = types.StringValue(dataItem.ProviderAccountID)
			data1.UpdatedAt = types.StringValue(dataItem.UpdatedAt.Format(time.RFC3339Nano))
			if dataCount+1 > len(r.Data) {
				r.Data = append(r.Data, data1)
			} else {
				r.Data[dataCount].CreatedAt = data1.CreatedAt
				r.Data[dataCount].ID = data1.ID
				r.Data[dataCount].Provider = data1.Provider
				r.Data[dataCount].ProviderAccountID = data1.ProviderAccountID
				r.Data[dataCount].UpdatedAt = data1.UpdatedAt
			}
		}
		r.Meta.Page.Number = types.NumberValue(big.NewFloat(float64(resp.Meta.Page.Number)))
		r.Meta.Page.Size = types.NumberValue(big.NewFloat(float64(resp.Meta.Page.Size)))
		r.Meta.Page.Total = types.NumberValue(big.NewFloat(float64(resp.Meta.Page.Total)))
	}
}
