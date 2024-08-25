// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
	speakeasy_int64validators "github.com/kong/terraform-provider-konnect/internal/validators/int64validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-konnect/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-konnect/internal/validators/stringvalidators"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudGatewayConfigurationResource{}
var _ resource.ResourceWithImportState = &CloudGatewayConfigurationResource{}

func NewCloudGatewayConfigurationResource() resource.Resource {
	return &CloudGatewayConfigurationResource{}
}

// CloudGatewayConfigurationResource defines the resource implementation.
type CloudGatewayConfigurationResource struct {
	client *sdk.Konnect
}

// CloudGatewayConfigurationResourceModel describes the resource data model.
type CloudGatewayConfigurationResourceModel struct {
	APIAccess            types.String                                `tfsdk:"api_access"`
	ControlPlaneGeo      types.String                                `tfsdk:"control_plane_geo"`
	ControlPlaneID       types.String                                `tfsdk:"control_plane_id"`
	CreatedAt            types.String                                `tfsdk:"created_at"`
	DataplaneGroupConfig []tfTypes.ConfigurationDataPlaneGroupConfig `tfsdk:"dataplane_group_config"`
	DataplaneGroups      []tfTypes.ConfigurationDataPlaneGroup       `tfsdk:"dataplane_groups"`
	EntityVersion        types.Number                                `tfsdk:"entity_version"`
	ID                   types.String                                `tfsdk:"id"`
	UpdatedAt            types.String                                `tfsdk:"updated_at"`
	Version              types.String                                `tfsdk:"version"`
}

func (r *CloudGatewayConfigurationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_gateway_configuration"
}

func (r *CloudGatewayConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CloudGatewayConfiguration Resource",
		Attributes: map[string]schema.Attribute{
			"api_access": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("private+public"),
				Description: `Type of API access data-plane groups will support for a configuration. must be one of ["private", "public", "private+public"]; Default: "private+public"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"private",
						"public",
						"private+public",
					),
				},
			},
			"control_plane_geo": schema.StringAttribute{
				Required:    true,
				Description: `Set of control-plane geos supported for deploying cloud-gateways configurations. must be one of ["us", "eu", "au"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"us",
						"eu",
						"au",
					),
				},
			},
			"control_plane_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of configuration creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"dataplane_group_config": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"autoscale": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"configuration_data_plane_group_autoscale_autopilot": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"base_rps": schema.Int64Attribute{
											Computed:    true,
											Description: `Base number of requests per second that the deployment target should support.`,
										},
										"kind": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["autopilot"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"autopilot",
												),
											},
										},
										"max_rps": schema.Int64Attribute{
											Computed:    true,
											Description: `Max number of requests per second that the deployment target should support. If not set, this defaults to 10x base_rps.`,
										},
									},
									Description: `Object that describes the autopilot autoscaling strategy.`,
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(path.Expressions{
											path.MatchRelative().AtParent().AtName("configuration_data_plane_group_autoscale_static"),
										}...),
									},
								},
								"configuration_data_plane_group_autoscale_static": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"instance_type": schema.StringAttribute{
											Computed:    true,
											Description: `Instance type name to indicate capacity. must be one of ["small", "medium", "large"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"small",
													"medium",
													"large",
												),
											},
										},
										"kind": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["static"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"static",
												),
											},
										},
										"requested_instances": schema.Int64Attribute{
											Computed:    true,
											Description: `Number of data-planes the deployment target will contain.`,
										},
									},
									Description: `Object that describes the static autoscaling strategy.`,
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(path.Expressions{
											path.MatchRelative().AtParent().AtName("configuration_data_plane_group_autoscale_autopilot"),
										}...),
									},
								},
							},
							Validators: []validator.Object{
								validators.ExactlyOneChild(),
							},
						},
						"cloud_gateway_network_id": schema.StringAttribute{
							Computed: true,
						},
						"provider": schema.StringAttribute{
							Computed:    true,
							Description: `Name of cloud provider. must be one of ["aws"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"aws",
								),
							},
						},
						"region": schema.StringAttribute{
							Computed:    true,
							Description: `Region ID for cloud provider region.`,
						},
					},
				},
				Description: `Object that describes where data-planes will be deployed to, along with how many instances.`,
			},
			"dataplane_groups": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"autoscale": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"configuration_data_plane_group_autoscale_autopilot": schema.SingleNestedAttribute{
									Computed: true,
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"base_rps": schema.Int64Attribute{
											Computed:    true,
											Optional:    true,
											Description: `Base number of requests per second that the deployment target should support. Not Null`,
											Validators: []validator.Int64{
												speakeasy_int64validators.NotNull(),
											},
										},
										"kind": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Not Null; must be one of ["autopilot"]`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
												stringvalidator.OneOf(
													"autopilot",
												),
											},
										},
										"max_rps": schema.Int64Attribute{
											Computed:    true,
											Optional:    true,
											Description: `Max number of requests per second that the deployment target should support. If not set, this defaults to 10x base_rps.`,
										},
									},
									Description: `Object that describes the autopilot autoscaling strategy.`,
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(path.Expressions{
											path.MatchRelative().AtParent().AtName("configuration_data_plane_group_autoscale_static"),
										}...),
									},
								},
								"configuration_data_plane_group_autoscale_static": schema.SingleNestedAttribute{
									Computed: true,
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"instance_type": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Instance type name to indicate capacity. Not Null; must be one of ["small", "medium", "large"]`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
												stringvalidator.OneOf(
													"small",
													"medium",
													"large",
												),
											},
										},
										"kind": schema.StringAttribute{
											Computed:    true,
											Optional:    true,
											Description: `Not Null; must be one of ["static"]`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
												stringvalidator.OneOf(
													"static",
												),
											},
										},
										"requested_instances": schema.Int64Attribute{
											Computed:    true,
											Optional:    true,
											Description: `Number of data-planes the deployment target will contain. Not Null`,
											Validators: []validator.Int64{
												speakeasy_int64validators.NotNull(),
											},
										},
									},
									Description: `Object that describes the static autoscaling strategy.`,
									Validators: []validator.Object{
										objectvalidator.ConflictsWith(path.Expressions{
											path.MatchRelative().AtParent().AtName("configuration_data_plane_group_autoscale_autopilot"),
										}...),
									},
								},
							},
							Description: `Not Null`,
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
								validators.ExactlyOneChild(),
							},
						},
						"cloud_gateway_network_id": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: `An RFC-3339 timestamp representation of data-plane group creation date.`,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
						"egress_ip_addresses": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							MarkdownDescription: `List of egress IP addresses for the network that this data-plane group runs on.` + "\n" +
								``,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `ID of the data-plane group that represents a deployment target for a set of data-planes.`,
						},
						"private_ip_addresses": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							MarkdownDescription: `List of private IP addresses of the internal load balancer that proxies traffic to this data-plane group.` + "\n" +
								``,
						},
						"provider": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Name of cloud provider. Not Null; must be one of ["aws"]`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
								stringvalidator.OneOf(
									"aws",
								),
							},
						},
						"region": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Region ID for cloud provider region. Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"state": schema.StringAttribute{
							Computed:    true,
							Description: `State of the data-plane group. must be one of ["created", "initializing", "ready", "terminating", "terminated"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"created",
									"initializing",
									"ready",
									"terminating",
									"terminated",
								),
							},
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: `An RFC-3339 timestamp representation of data-plane group update date.`,
							Validators: []validator.String{
								validators.IsRFC3339(),
							},
						},
					},
				},
				Description: `List of data-plane groups that describe where to deploy instances, along with how many instances.`,
			},
			"entity_version": schema.NumberAttribute{
				Computed: true,
				MarkdownDescription: `Positive, monotonically increasing version integer, to serialize configuration changes.` + "\n" +
					``,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of configuration update date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"version": schema.StringAttribute{
				Required:    true,
				Description: `Supported gateway version.`,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^(\d)+(\.(\d)+)?$`), "must match pattern "+regexp.MustCompile(`^(\d)+(\.(\d)+)?$`).String()),
				},
			},
		},
	}
}

func (r *CloudGatewayConfigurationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CloudGatewayConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CloudGatewayConfigurationResourceModel
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

	request := *data.ToSharedCreateConfigurationRequest()
	res, err := r.client.DataPlaneGroupConfigurations.CreateConfiguration(ctx, request)
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
	if !(res.ConfigurationManifest != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConfigurationManifest(res.ConfigurationManifest)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CloudGatewayConfigurationResourceModel
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

	var configurationID string
	configurationID = data.ID.ValueString()

	request := operations.GetConfigurationRequest{
		ConfigurationID: configurationID,
	}
	res, err := r.client.DataPlaneGroupConfigurations.GetConfiguration(ctx, request)
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
	if !(res.ConfigurationManifest != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConfigurationManifest(res.ConfigurationManifest)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CloudGatewayConfigurationResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedCreateConfigurationRequest()
	res, err := r.client.DataPlaneGroupConfigurations.CreateConfiguration(ctx, request)
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
	if !(res.ConfigurationManifest != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConfigurationManifest(res.ConfigurationManifest)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CloudGatewayConfigurationResourceModel
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

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *CloudGatewayConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
