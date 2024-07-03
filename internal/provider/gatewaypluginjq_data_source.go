// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayPluginJQDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginJQDataSource{}

func NewGatewayPluginJQDataSource() datasource.DataSource {
	return &GatewayPluginJQDataSource{}
}

// GatewayPluginJQDataSource is the data source implementation.
type GatewayPluginJQDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginJQDataSourceModel describes the data model.
type GatewayPluginJQDataSourceModel struct {
	Config         *tfTypes.CreateJQPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer          `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLConsumer          `tfsdk:"consumer_group"`
	ControlPlaneID types.String                  `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                   `tfsdk:"created_at"`
	Enabled        types.Bool                    `tfsdk:"enabled"`
	ID             types.String                  `tfsdk:"id"`
	InstanceName   types.String                  `tfsdk:"instance_name"`
	Protocols      []types.String                `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer          `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer          `tfsdk:"service"`
	Tags           []types.String                `tfsdk:"tags"`
	UpdatedAt      types.Int64                   `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginJQDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_jq"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginJQDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginJQ DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"request_if_media_type": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"request_jq_program": schema.StringAttribute{
						Computed: true,
					},
					"request_jq_program_options": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"ascii_output": schema.BoolAttribute{
								Computed: true,
							},
							"compact_output": schema.BoolAttribute{
								Computed: true,
							},
							"join_output": schema.BoolAttribute{
								Computed: true,
							},
							"raw_output": schema.BoolAttribute{
								Computed: true,
							},
							"sort_keys": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"response_if_media_type": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"response_if_status_code": schema.ListAttribute{
						Computed:    true,
						ElementType: types.Int64Type,
					},
					"response_jq_program": schema.StringAttribute{
						Computed: true,
					},
					"response_jq_program_options": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"ascii_output": schema.BoolAttribute{
								Computed: true,
							},
							"compact_output": schema.BoolAttribute{
								Computed: true,
							},
							"join_output": schema.BoolAttribute{
								Computed: true,
							},
							"raw_output": schema.BoolAttribute{
								Computed: true,
							},
							"sort_keys": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager.`,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayPluginJQDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayPluginJQDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginJQDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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

	pluginID := data.ID.ValueString()
	controlPlaneID := data.ControlPlaneID.ValueString()
	request := operations.GetJqPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetJqPlugin(ctx, request)
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
	if !(res.JQPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedJQPlugin(res.JQPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
