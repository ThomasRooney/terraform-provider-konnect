// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"time"
)

func (r *CloudGatewayNetworkResourceModel) ToSharedCreateNetworkRequest() *shared.CreateNetworkRequest {
	var name string
	name = r.Name.ValueString()

	var cloudGatewayProviderAccountID string
	cloudGatewayProviderAccountID = r.CloudGatewayProviderAccountID.ValueString()

	var region string
	region = r.Region.ValueString()

	var availabilityZones []string = []string{}
	for _, availabilityZonesItem := range r.AvailabilityZones {
		availabilityZones = append(availabilityZones, availabilityZonesItem.ValueString())
	}
	var cidrBlock string
	cidrBlock = r.CidrBlock.ValueString()

	var firewall *shared.NetworkFirewallConfig
	if r.Firewall != nil {
		var allowedCidrBlocks []string = []string{}
		for _, allowedCidrBlocksItem := range r.Firewall.AllowedCidrBlocks {
			allowedCidrBlocks = append(allowedCidrBlocks, allowedCidrBlocksItem.ValueString())
		}
		var deniedCidrBlocks []string = []string{}
		for _, deniedCidrBlocksItem := range r.Firewall.DeniedCidrBlocks {
			deniedCidrBlocks = append(deniedCidrBlocks, deniedCidrBlocksItem.ValueString())
		}
		firewall = &shared.NetworkFirewallConfig{
			AllowedCidrBlocks: allowedCidrBlocks,
			DeniedCidrBlocks:  deniedCidrBlocks,
		}
	}
	ddosProtection := new(bool)
	if !r.DdosProtection.IsUnknown() && !r.DdosProtection.IsNull() {
		*ddosProtection = r.DdosProtection.ValueBool()
	} else {
		ddosProtection = nil
	}
	out := shared.CreateNetworkRequest{
		Name:                          name,
		CloudGatewayProviderAccountID: cloudGatewayProviderAccountID,
		Region:                        region,
		AvailabilityZones:             availabilityZones,
		CidrBlock:                     cidrBlock,
		Firewall:                      firewall,
		DdosProtection:                ddosProtection,
	}
	return &out
}

func (r *CloudGatewayNetworkResourceModel) RefreshFromSharedNetwork(resp *shared.Network) {
	if resp != nil {
		r.AvailabilityZones = []types.String{}
		for _, v := range resp.AvailabilityZones {
			r.AvailabilityZones = append(r.AvailabilityZones, types.StringValue(v))
		}
		r.CidrBlock = types.StringValue(resp.CidrBlock)
		r.CloudGatewayProviderAccountID = types.StringValue(resp.CloudGatewayProviderAccountID)
		r.ConfigurationReferenceCount = types.Int64Value(resp.ConfigurationReferenceCount)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		r.DdosProtection = types.BoolPointerValue(resp.DdosProtection)
		r.Default = types.BoolValue(resp.Default)
		r.EntityVersion = types.Int64Value(resp.EntityVersion)
		if resp.Firewall == nil {
			r.Firewall = nil
		} else {
			r.Firewall = &tfTypes.NetworkFirewallConfig{}
			r.Firewall.AllowedCidrBlocks = []types.String{}
			for _, v := range resp.Firewall.AllowedCidrBlocks {
				r.Firewall.AllowedCidrBlocks = append(r.Firewall.AllowedCidrBlocks, types.StringValue(v))
			}
			r.Firewall.DeniedCidrBlocks = []types.String{}
			for _, v := range resp.Firewall.DeniedCidrBlocks {
				r.Firewall.DeniedCidrBlocks = append(r.Firewall.DeniedCidrBlocks, types.StringValue(v))
			}
		}
		r.ID = types.StringValue(resp.ID)
		r.Name = types.StringValue(resp.Name)
		r.ProviderMetadata.SubnetIds = []types.String{}
		for _, v := range resp.ProviderMetadata.SubnetIds {
			r.ProviderMetadata.SubnetIds = append(r.ProviderMetadata.SubnetIds, types.StringValue(v))
		}
		r.ProviderMetadata.VpcID = types.StringPointerValue(resp.ProviderMetadata.VpcID)
		r.Region = types.StringValue(resp.Region)
		r.State = types.StringValue(string(resp.State))
		r.TransitGatewayCount = types.Int64Value(resp.TransitGatewayCount)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
	}
}

func (r *CloudGatewayNetworkResourceModel) ToSharedPatchNetworkRequest() *shared.PatchNetworkRequest {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	var firewall *shared.NetworkFirewallConfig
	if r.Firewall != nil {
		var allowedCidrBlocks []string = []string{}
		for _, allowedCidrBlocksItem := range r.Firewall.AllowedCidrBlocks {
			allowedCidrBlocks = append(allowedCidrBlocks, allowedCidrBlocksItem.ValueString())
		}
		var deniedCidrBlocks []string = []string{}
		for _, deniedCidrBlocksItem := range r.Firewall.DeniedCidrBlocks {
			deniedCidrBlocks = append(deniedCidrBlocks, deniedCidrBlocksItem.ValueString())
		}
		firewall = &shared.NetworkFirewallConfig{
			AllowedCidrBlocks: allowedCidrBlocks,
			DeniedCidrBlocks:  deniedCidrBlocks,
		}
	}
	out := shared.PatchNetworkRequest{
		Name:     name,
		Firewall: firewall,
	}
	return &out
}
