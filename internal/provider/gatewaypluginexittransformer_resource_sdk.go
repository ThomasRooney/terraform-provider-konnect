// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginExitTransformerResourceModel) ToSharedCreateExitTransformerPlugin() *shared.CreateExitTransformerPlugin {
	var config *shared.CreateExitTransformerPluginConfig
	if r.Config != nil {
		var functions []string = []string{}
		for _, functionsItem := range r.Config.Functions {
			functions = append(functions, functionsItem.ValueString())
		}
		handleUnexpected := new(bool)
		if !r.Config.HandleUnexpected.IsUnknown() && !r.Config.HandleUnexpected.IsNull() {
			*handleUnexpected = r.Config.HandleUnexpected.ValueBool()
		} else {
			handleUnexpected = nil
		}
		handleUnknown := new(bool)
		if !r.Config.HandleUnknown.IsUnknown() && !r.Config.HandleUnknown.IsNull() {
			*handleUnknown = r.Config.HandleUnknown.ValueBool()
		} else {
			handleUnknown = nil
		}
		config = &shared.CreateExitTransformerPluginConfig{
			Functions:        functions,
			HandleUnexpected: handleUnexpected,
			HandleUnknown:    handleUnknown,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var protocols []shared.CreateExitTransformerPluginProtocols = []shared.CreateExitTransformerPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateExitTransformerPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateExitTransformerPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateExitTransformerPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateExitTransformerPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateExitTransformerPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateExitTransformerPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateExitTransformerPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateExitTransformerPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateExitTransformerPluginService{
			ID: id3,
		}
	}
	out := shared.CreateExitTransformerPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginExitTransformerResourceModel) RefreshFromSharedExitTransformerPlugin(resp *shared.ExitTransformerPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateExitTransformerPluginConfig{}
			r.Config.Functions = []types.String{}
			for _, v := range resp.Config.Functions {
				r.Config.Functions = append(r.Config.Functions, types.StringValue(v))
			}
			r.Config.HandleUnexpected = types.BoolPointerValue(resp.Config.HandleUnexpected)
			r.Config.HandleUnknown = types.BoolPointerValue(resp.Config.HandleUnknown)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
