// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginJQDataSourceModel) RefreshFromSharedJQPlugin(resp *shared.JQPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateJQPluginConfig{}
			r.Config.RequestIfMediaType = []types.String{}
			for _, v := range resp.Config.RequestIfMediaType {
				r.Config.RequestIfMediaType = append(r.Config.RequestIfMediaType, types.StringValue(v))
			}
			r.Config.RequestJqProgram = types.StringPointerValue(resp.Config.RequestJqProgram)
			if resp.Config.RequestJqProgramOptions == nil {
				r.Config.RequestJqProgramOptions = nil
			} else {
				r.Config.RequestJqProgramOptions = &tfTypes.CreateJQPluginRequestJQProgramOptions{}
				r.Config.RequestJqProgramOptions.ASCIIOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.ASCIIOutput)
				r.Config.RequestJqProgramOptions.CompactOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.CompactOutput)
				r.Config.RequestJqProgramOptions.JoinOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.JoinOutput)
				r.Config.RequestJqProgramOptions.RawOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.RawOutput)
				r.Config.RequestJqProgramOptions.SortKeys = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.SortKeys)
			}
			r.Config.ResponseIfMediaType = []types.String{}
			for _, v := range resp.Config.ResponseIfMediaType {
				r.Config.ResponseIfMediaType = append(r.Config.ResponseIfMediaType, types.StringValue(v))
			}
			r.Config.ResponseIfStatusCode = []types.Int64{}
			for _, v := range resp.Config.ResponseIfStatusCode {
				r.Config.ResponseIfStatusCode = append(r.Config.ResponseIfStatusCode, types.Int64Value(v))
			}
			r.Config.ResponseJqProgram = types.StringPointerValue(resp.Config.ResponseJqProgram)
			if resp.Config.ResponseJqProgramOptions == nil {
				r.Config.ResponseJqProgramOptions = nil
			} else {
				r.Config.ResponseJqProgramOptions = &tfTypes.CreateJQPluginRequestJQProgramOptions{}
				r.Config.ResponseJqProgramOptions.ASCIIOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.ASCIIOutput)
				r.Config.ResponseJqProgramOptions.CompactOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.CompactOutput)
				r.Config.ResponseJqProgramOptions.JoinOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.JoinOutput)
				r.Config.ResponseJqProgramOptions.RawOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.RawOutput)
				r.Config.ResponseJqProgramOptions.SortKeys = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.SortKeys)
			}
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
