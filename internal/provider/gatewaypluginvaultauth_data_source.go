// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

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
var _ datasource.DataSource = &GatewayPluginVaultAuthDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginVaultAuthDataSource{}

func NewGatewayPluginVaultAuthDataSource() datasource.DataSource {
	return &GatewayPluginVaultAuthDataSource{}
}

// GatewayPluginVaultAuthDataSource is the data source implementation.
type GatewayPluginVaultAuthDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginVaultAuthDataSourceModel describes the data model.
type GatewayPluginVaultAuthDataSourceModel struct {
	Config         *tfTypes.CreateVaultAuthPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                 `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLConsumer                 `tfsdk:"consumer_group"`
	ControlPlaneID types.String                         `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                          `tfsdk:"created_at"`
	Enabled        types.Bool                           `tfsdk:"enabled"`
	ID             types.String                         `tfsdk:"id"`
	InstanceName   types.String                         `tfsdk:"instance_name"`
	Ordering       *tfTypes.CreateACLPluginOrdering     `tfsdk:"ordering"`
	Protocols      []types.String                       `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                 `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                 `tfsdk:"service"`
	Tags           []types.String                       `tfsdk:"tags"`
	UpdatedAt      types.Int64                          `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginVaultAuthDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_vault_auth"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginVaultAuthDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginVaultAuth DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"access_token_name": schema.StringAttribute{
						Computed:    true,
						Description: `Describes an array of comma-separated parameter names where the plugin looks for an access token. The client must send the access token in one of those key names, and the plugin will try to read the credential from a header or the querystring parameter with the same name. The key names can only contain [a-z], [A-Z], [0-9], [_], and [-].`,
					},
					"anonymous": schema.StringAttribute{
						Computed:    true,
						Description: `An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request fails with an authentication failure ` + "`" + `4xx` + "`" + `. Note that this value must refer to the consumer ` + "`" + `id` + "`" + ` or ` + "`" + `username` + "`" + ` attribute, and **not** its ` + "`" + `custom_id` + "`" + `.`,
					},
					"hide_credentials": schema.BoolAttribute{
						Computed:    true,
						Description: `An optional boolean value telling the plugin to show or hide the credential from the upstream service. If ` + "`" + `true` + "`" + `, the plugin will strip the credential from the request (i.e. the header or querystring containing the key) before proxying it.`,
					},
					"run_on_preflight": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that indicates whether the plugin should run (and try to authenticate) on ` + "`" + `OPTIONS` + "`" + ` preflight requests. If set to ` + "`" + `false` + "`" + `, then ` + "`" + `OPTIONS` + "`" + ` requests will always be allowed.`,
					},
					"secret_token_name": schema.StringAttribute{
						Computed:    true,
						Description: `Describes an array of comma-separated parameter names where the plugin looks for a secret token. The client must send the secret in one of those key names, and the plugin will try to read the credential from a header or the querystring parameter with the same name. The key names can only contain [a-z], [A-Z], [0-9], [_], and [-].`,
					},
					"tokens_in_body": schema.BoolAttribute{
						Computed:    true,
						Description: `If enabled, the plugin will read the request body (if said request has one and its MIME type is supported) and try to find the key in it. Supported MIME types are ` + "`" + `application/www-form-urlencoded` + "`" + `, ` + "`" + `application/json` + "`" + `, and ` + "`" + `multipart/form-data` + "`" + `.`,
					},
					"vault": schema.StringAttribute{
						Computed:    true,
						Description: `A reference to an existing ` + "`" + `vault` + "`" + ` object within the database. ` + "`" + `vault` + "`" + ` entities define the connection and authentication parameters used to connect to a Vault HTTP(S) API.`,
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
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed. `,
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
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
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

func (r *GatewayPluginVaultAuthDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GatewayPluginVaultAuthDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginVaultAuthDataSourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetVaultauthPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetVaultauthPlugin(ctx, request)
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
	if !(res.VaultAuthPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedVaultAuthPlugin(res.VaultAuthPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}