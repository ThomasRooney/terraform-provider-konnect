// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayPluginXMLThreatProtectionResource{}
var _ resource.ResourceWithImportState = &GatewayPluginXMLThreatProtectionResource{}

func NewGatewayPluginXMLThreatProtectionResource() resource.Resource {
	return &GatewayPluginXMLThreatProtectionResource{}
}

// GatewayPluginXMLThreatProtectionResource defines the resource implementation.
type GatewayPluginXMLThreatProtectionResource struct {
	client *sdk.Konnect
}

// GatewayPluginXMLThreatProtectionResourceModel describes the resource data model.
type GatewayPluginXMLThreatProtectionResourceModel struct {
	Config         *tfTypes.CreateXMLThreatProtectionPluginConfig `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer                           `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLConsumer                           `tfsdk:"consumer_group"`
	ControlPlaneID types.String                                   `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                                    `tfsdk:"created_at"`
	Enabled        types.Bool                                     `tfsdk:"enabled"`
	ID             types.String                                   `tfsdk:"id"`
	InstanceName   types.String                                   `tfsdk:"instance_name"`
	Ordering       *tfTypes.CreateACLPluginOrdering               `tfsdk:"ordering"`
	Protocols      []types.String                                 `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer                           `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer                           `tfsdk:"service"`
	Tags           []types.String                                 `tfsdk:"tags"`
	UpdatedAt      types.Int64                                    `tfsdk:"updated_at"`
}

func (r *GatewayPluginXMLThreatProtectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_xml_threat_protection"
}

func (r *GatewayPluginXMLThreatProtectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginXMLThreatProtection Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"allow_dtd": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Indicates whether an XML Document Type Definition (DTD) section is allowed.`,
					},
					"allowed_content_types": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `A list of Content-Type values with payloads that are allowed, but aren't validated.`,
					},
					"attribute": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of the attribute value.`,
					},
					"bla_max_amplification": schema.NumberAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Sets the maximum allowed amplification. This protects against the Billion Laughs Attack.`,
					},
					"bla_threshold": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Sets the threshold after which the protection starts. This protects against the Billion Laughs Attack.`,
						Validators: []validator.Int64{
							int64validator.AtLeast(1024),
						},
					},
					"buffer": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of the unparsed buffer (see below).`,
					},
					"checked_content_types": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `A list of Content-Type values with payloads that must be validated.`,
					},
					"comment": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of comments.`,
					},
					"document": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of the entire document.`,
					},
					"entity": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of entity values in EntityDecl.`,
					},
					"entityname": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of entity names in EntityDecl.`,
					},
					"entityproperty": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of systemId, publicId, or notationName in EntityDecl.`,
					},
					"localname": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of the localname. This applies to tags and attributes.`,
					},
					"max_attributes": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum number of attributes allowed on a tag, including default ones. Note: If namespace-aware parsing is disabled, then the namespaces definitions are counted as attributes.`,
					},
					"max_children": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum number of children allowed (Element, Text, Comment, ProcessingInstruction, CDATASection). Note: Adjacent text and CDATA sections are counted as one. For example, text-cdata-text-cdata is one child.`,
					},
					"max_depth": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum depth of tags. Child elements such as Text or Comments are not counted as another level.`,
					},
					"max_namespaces": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum number of namespaces defined on a tag. This value is required if parsing is namespace-aware.`,
					},
					"namespace_aware": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `If not parsing namespace aware, all prefixes and namespace attributes will be counted as regular attributes and element names, and validated as such.`,
					},
					"namespaceuri": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of the namespace URI. This value is required if parsing is namespace-aware.`,
					},
					"pidata": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of processing instruction data.`,
					},
					"pitarget": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of processing instruction targets.`,
					},
					"prefix": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum size of the prefix. This applies to tags and attributes. This value is required if parsing is namespace-aware.`,
					},
					"text": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Description: `Maximum text inside tags (counted over all adjacent text/CDATA elements combined).`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"control_plane_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed. `,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
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

func (r *GatewayPluginXMLThreatProtectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayPluginXMLThreatProtectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayPluginXMLThreatProtectionResourceModel
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

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	createXMLThreatProtectionPlugin := data.ToSharedCreateXMLThreatProtectionPlugin()
	request := operations.CreateXmlthreatprotectionPluginRequest{
		ControlPlaneID:                  controlPlaneID,
		CreateXMLThreatProtectionPlugin: createXMLThreatProtectionPlugin,
	}
	res, err := r.client.Plugins.CreateXmlthreatprotectionPlugin(ctx, request)
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
	if !(res.XMLThreatProtectionPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedXMLThreatProtectionPlugin(res.XMLThreatProtectionPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginXMLThreatProtectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayPluginXMLThreatProtectionResourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetXmlthreatprotectionPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetXmlthreatprotectionPlugin(ctx, request)
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
	if !(res.XMLThreatProtectionPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedXMLThreatProtectionPlugin(res.XMLThreatProtectionPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginXMLThreatProtectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayPluginXMLThreatProtectionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	createXMLThreatProtectionPlugin := data.ToSharedCreateXMLThreatProtectionPlugin()
	request := operations.UpdateXmlthreatprotectionPluginRequest{
		PluginID:                        pluginID,
		ControlPlaneID:                  controlPlaneID,
		CreateXMLThreatProtectionPlugin: createXMLThreatProtectionPlugin,
	}
	res, err := r.client.Plugins.UpdateXmlthreatprotectionPlugin(ctx, request)
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
	if !(res.XMLThreatProtectionPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedXMLThreatProtectionPlugin(res.XMLThreatProtectionPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayPluginXMLThreatProtectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayPluginXMLThreatProtectionResourceModel
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

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.DeleteXmlthreatprotectionPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.DeleteXmlthreatprotectionPlugin(ctx, request)
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

func (r *GatewayPluginXMLThreatProtectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		ControlPlaneID string `json:"control_plane_id"`
		ID             string `json:"id"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The ID is not valid. It's expected to be a JSON object alike '{ "control_plane_id": "9524ec7d-36d9-465d-a8c5-83a3c9390458",  "plugin_id": "3473c251-5b6c-4f45-b1ff-7ede735a366d"}': `+err.Error())
		return
	}

	if len(data.ControlPlaneID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field control_plane_id is required but was not found in the json encoded ID. It's expected to be a value alike '"9524ec7d-36d9-465d-a8c5-83a3c9390458"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("control_plane_id"), data.ControlPlaneID)...)
	if len(data.ID) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field id is required but was not found in the json encoded ID. It's expected to be a value alike '"3473c251-5b6c-4f45-b1ff-7ede735a366d"`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.ID)...)

}