// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_boolplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/boolplanmodifier"
	speakeasy_listplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/listplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudGatewayNetworkResource{}
var _ resource.ResourceWithImportState = &CloudGatewayNetworkResource{}

func NewCloudGatewayNetworkResource() resource.Resource {
	return &CloudGatewayNetworkResource{}
}

// CloudGatewayNetworkResource defines the resource implementation.
type CloudGatewayNetworkResource struct {
	client *sdk.Konnect
}

// CloudGatewayNetworkResourceModel describes the resource data model.
type CloudGatewayNetworkResourceModel struct {
	AvailabilityZones             []types.String                  `tfsdk:"availability_zones"`
	CidrBlock                     types.String                    `tfsdk:"cidr_block"`
	CloudGatewayProviderAccountID types.String                    `tfsdk:"cloud_gateway_provider_account_id"`
	ConfigurationReferenceCount   types.Int64                     `tfsdk:"configuration_reference_count"`
	CreatedAt                     types.String                    `tfsdk:"created_at"`
	DdosProtection                types.Bool                      `tfsdk:"ddos_protection"`
	Default                       types.Bool                      `tfsdk:"default"`
	EntityVersion                 types.Int64                     `tfsdk:"entity_version"`
	Firewall                      *tfTypes.NetworkFirewallConfig  `tfsdk:"firewall"`
	ID                            types.String                    `tfsdk:"id"`
	Name                          types.String                    `tfsdk:"name"`
	ProviderMetadata              tfTypes.NetworkProviderMetadata `tfsdk:"provider_metadata"`
	Region                        types.String                    `tfsdk:"region"`
	State                         types.String                    `tfsdk:"state"`
	TransitGatewayCount           types.Int64                     `tfsdk:"transit_gateway_count"`
	UpdatedAt                     types.String                    `tfsdk:"updated_at"`
}

func (r *CloudGatewayNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_gateway_network"
}

func (r *CloudGatewayNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CloudGatewayNetwork Resource",
		Attributes: map[string]schema.Attribute{
			"availability_zones": schema.ListAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				ElementType: types.StringType,
				Description: `List of availability zones that the network is attached to. Requires replacement if changed. `,
			},
			"cidr_block": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `CIDR block configuration for the network. Requires replacement if changed. `,
			},
			"cloud_gateway_provider_account_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Requires replacement if changed. `,
			},
			"configuration_reference_count": schema.Int64Attribute{
				Computed:    true,
				Description: `The number of configurations that reference this network.`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of network creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"ddos_protection": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
				},
				Optional:    true,
				Description: `Whether DDOS protection is enabled for the network. Requires replacement if changed. `,
			},
			"default": schema.BoolAttribute{
				Computed: true,
				MarkdownDescription: `Whether the network is a default network or not. Default networks are Networks that are created` + "\n" +
					`automatically by Konnect when an organization is linked to a provider account.` + "\n" +
					``,
			},
			"entity_version": schema.Int64Attribute{
				Computed: true,
				MarkdownDescription: `Monotonically-increasing version count of the network, to indicate the order of updates to the network.` + "\n" +
					``,
			},
			"firewall": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"allowed_cidr_blocks": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of allowed CIDR blocks to access a network.`,
					},
					"denied_cidr_blocks": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of denied CIDR blocks to access a network.`,
					},
				},
				Description: `Firewall configuration for a network.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The network to operate on.`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `Human-readable name of the network.`,
			},
			"provider_metadata": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"subnet_ids": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"vpc_id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `Metadata describing attributes returned by cloud-provider for the network.`,
			},
			"region": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Region ID for cloud provider region. Requires replacement if changed. `,
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: `State of the network. must be one of ["created", "initializing", "offline", "ready", "terminating", "terminated"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"created",
						"initializing",
						"offline",
						"ready",
						"terminating",
						"terminated",
					),
				},
			},
			"transit_gateway_count": schema.Int64Attribute{
				Computed:    true,
				Description: `The number of transit gateways attached to this network.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of network update date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *CloudGatewayNetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *CloudGatewayNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CloudGatewayNetworkResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedCreateNetworkRequest()
	res, err := r.client.Networks.CreateNetwork(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Network == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedNetwork(res.Network)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CloudGatewayNetworkResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	networkID := data.ID.ValueString()
	request := operations.GetNetworkRequest{
		NetworkID: networkID,
	}
	res, err := r.client.Networks.GetNetwork(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Network == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedNetwork(res.Network)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CloudGatewayNetworkResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	networkID := data.ID.ValueString()
	patchNetworkRequest := *data.ToSharedPatchNetworkRequest()
	request := operations.UpdateNetworkRequest{
		NetworkID:           networkID,
		PatchNetworkRequest: patchNetworkRequest,
	}
	res, err := r.client.Networks.UpdateNetwork(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.Network == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedNetwork(res.Network)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CloudGatewayNetworkResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	networkID := data.ID.ValueString()
	request := operations.DeleteNetworkRequest{
		NetworkID: networkID,
	}
	res, err := r.client.Networks.DeleteNetwork(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *CloudGatewayNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}