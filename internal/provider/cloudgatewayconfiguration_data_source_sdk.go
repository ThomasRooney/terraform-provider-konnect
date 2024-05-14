// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *CloudGatewayConfigurationDataSourceModel) RefreshFromSharedConfigurationManifest(resp *shared.ConfigurationManifest) {
	if resp != nil {
		if resp.APIAccess != nil {
			r.APIAccess = types.StringValue(string(*resp.APIAccess))
		} else {
			r.APIAccess = types.StringNull()
		}
		r.ControlPlaneGeo = types.StringValue(string(resp.ControlPlaneGeo))
		r.ControlPlaneID = types.StringValue(resp.ControlPlaneID)
		r.CreatedAt = types.StringValue(resp.CreatedAt.Format(time.RFC3339Nano))
		if len(r.DataplaneGroupConfig) > len(resp.DataplaneGroupConfig) {
			r.DataplaneGroupConfig = r.DataplaneGroupConfig[:len(resp.DataplaneGroupConfig)]
		}
		for dataplaneGroupConfigCount, dataplaneGroupConfigItem := range resp.DataplaneGroupConfig {
			var dataplaneGroupConfig1 tfTypes.ConfigurationDataPlaneGroupConfig
			if dataplaneGroupConfigItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot != nil {
				dataplaneGroupConfig1.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot = &tfTypes.ConfigurationDataPlaneGroupAutoscaleAutopilot{}
				dataplaneGroupConfig1.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.BaseRps = types.Int64Value(dataplaneGroupConfigItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.BaseRps)
				dataplaneGroupConfig1.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.Kind = types.StringValue(string(dataplaneGroupConfigItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.Kind))
				dataplaneGroupConfig1.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.MaxRps = types.Int64PointerValue(dataplaneGroupConfigItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.MaxRps)
			}
			if dataplaneGroupConfigItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic != nil {
				dataplaneGroupConfig1.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic = &tfTypes.ConfigurationDataPlaneGroupAutoscaleStatic{}
				dataplaneGroupConfig1.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.InstanceType = types.StringValue(string(dataplaneGroupConfigItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.InstanceType))
				dataplaneGroupConfig1.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.Kind = types.StringValue(string(dataplaneGroupConfigItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.Kind))
				dataplaneGroupConfig1.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.RequestedInstances = types.Int64Value(dataplaneGroupConfigItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.RequestedInstances)
			}
			dataplaneGroupConfig1.CloudGatewayNetworkID = types.StringValue(dataplaneGroupConfigItem.CloudGatewayNetworkID)
			dataplaneGroupConfig1.Provider = types.StringValue(string(dataplaneGroupConfigItem.Provider))
			dataplaneGroupConfig1.Region = types.StringValue(dataplaneGroupConfigItem.Region)
			if dataplaneGroupConfigCount+1 > len(r.DataplaneGroupConfig) {
				r.DataplaneGroupConfig = append(r.DataplaneGroupConfig, dataplaneGroupConfig1)
			} else {
				r.DataplaneGroupConfig[dataplaneGroupConfigCount].Autoscale = dataplaneGroupConfig1.Autoscale
				r.DataplaneGroupConfig[dataplaneGroupConfigCount].CloudGatewayNetworkID = dataplaneGroupConfig1.CloudGatewayNetworkID
				r.DataplaneGroupConfig[dataplaneGroupConfigCount].Provider = dataplaneGroupConfig1.Provider
				r.DataplaneGroupConfig[dataplaneGroupConfigCount].Region = dataplaneGroupConfig1.Region
			}
		}
		if len(r.DataplaneGroups) > len(resp.DataplaneGroups) {
			r.DataplaneGroups = r.DataplaneGroups[:len(resp.DataplaneGroups)]
		}
		for dataplaneGroupsCount, dataplaneGroupsItem := range resp.DataplaneGroups {
			var dataplaneGroups1 tfTypes.ConfigurationDataPlaneGroup
			if dataplaneGroupsItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot != nil {
				dataplaneGroups1.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot = &tfTypes.ConfigurationDataPlaneGroupAutoscaleAutopilot{}
				dataplaneGroups1.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.BaseRps = types.Int64Value(dataplaneGroupsItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.BaseRps)
				dataplaneGroups1.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.Kind = types.StringValue(string(dataplaneGroupsItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.Kind))
				dataplaneGroups1.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.MaxRps = types.Int64PointerValue(dataplaneGroupsItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleAutopilot.MaxRps)
			}
			if dataplaneGroupsItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic != nil {
				dataplaneGroups1.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic = &tfTypes.ConfigurationDataPlaneGroupAutoscaleStatic{}
				dataplaneGroups1.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.InstanceType = types.StringValue(string(dataplaneGroupsItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.InstanceType))
				dataplaneGroups1.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.Kind = types.StringValue(string(dataplaneGroupsItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.Kind))
				dataplaneGroups1.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.RequestedInstances = types.Int64Value(dataplaneGroupsItem.Autoscale.ConfigurationDataPlaneGroupAutoscaleStatic.RequestedInstances)
			}
			dataplaneGroups1.CloudGatewayNetworkID = types.StringValue(dataplaneGroupsItem.CloudGatewayNetworkID)
			dataplaneGroups1.CreatedAt = types.StringValue(dataplaneGroupsItem.CreatedAt.Format(time.RFC3339Nano))
			dataplaneGroups1.EgressIPAddresses = []types.String{}
			for _, v := range dataplaneGroupsItem.EgressIPAddresses {
				dataplaneGroups1.EgressIPAddresses = append(dataplaneGroups1.EgressIPAddresses, types.StringValue(v))
			}
			dataplaneGroups1.ID = types.StringValue(dataplaneGroupsItem.ID)
			dataplaneGroups1.PrivateIPAddresses = []types.String{}
			for _, v := range dataplaneGroupsItem.PrivateIPAddresses {
				dataplaneGroups1.PrivateIPAddresses = append(dataplaneGroups1.PrivateIPAddresses, types.StringValue(v))
			}
			dataplaneGroups1.Provider = types.StringValue(string(dataplaneGroupsItem.Provider))
			dataplaneGroups1.Region = types.StringValue(dataplaneGroupsItem.Region)
			dataplaneGroups1.State = types.StringValue(string(dataplaneGroupsItem.State))
			dataplaneGroups1.UpdatedAt = types.StringValue(dataplaneGroupsItem.UpdatedAt.Format(time.RFC3339Nano))
			if dataplaneGroupsCount+1 > len(r.DataplaneGroups) {
				r.DataplaneGroups = append(r.DataplaneGroups, dataplaneGroups1)
			} else {
				r.DataplaneGroups[dataplaneGroupsCount].Autoscale = dataplaneGroups1.Autoscale
				r.DataplaneGroups[dataplaneGroupsCount].CloudGatewayNetworkID = dataplaneGroups1.CloudGatewayNetworkID
				r.DataplaneGroups[dataplaneGroupsCount].CreatedAt = dataplaneGroups1.CreatedAt
				r.DataplaneGroups[dataplaneGroupsCount].EgressIPAddresses = dataplaneGroups1.EgressIPAddresses
				r.DataplaneGroups[dataplaneGroupsCount].ID = dataplaneGroups1.ID
				r.DataplaneGroups[dataplaneGroupsCount].PrivateIPAddresses = dataplaneGroups1.PrivateIPAddresses
				r.DataplaneGroups[dataplaneGroupsCount].Provider = dataplaneGroups1.Provider
				r.DataplaneGroups[dataplaneGroupsCount].Region = dataplaneGroups1.Region
				r.DataplaneGroups[dataplaneGroupsCount].State = dataplaneGroups1.State
				r.DataplaneGroups[dataplaneGroupsCount].UpdatedAt = dataplaneGroups1.UpdatedAt
			}
		}
		r.EntityVersion = types.NumberValue(big.NewFloat(float64(resp.EntityVersion)))
		r.ID = types.StringValue(resp.ID)
		r.UpdatedAt = types.StringValue(resp.UpdatedAt.Format(time.RFC3339Nano))
		r.Version = types.StringValue(resp.Version)
	}
}
