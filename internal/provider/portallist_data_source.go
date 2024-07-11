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
var _ datasource.DataSource = &PortalListDataSource{}
var _ datasource.DataSourceWithConfigure = &PortalListDataSource{}

func NewPortalListDataSource() datasource.DataSource {
	return &PortalListDataSource{}
}

// PortalListDataSource is the data source implementation.
type PortalListDataSource struct {
	client *sdk.Konnect
}

// PortalListDataSourceModel describes the data model.
type PortalListDataSourceModel struct {
	Data       []tfTypes.Portal      `tfsdk:"data"`
	Meta       tfTypes.PaginatedMeta `tfsdk:"meta"`
	PageNumber types.Int64           `tfsdk:"page_number"`
	PageSize   types.Int64           `tfsdk:"page_size"`
	Sort       types.String          `tfsdk:"sort"`
}

// Metadata returns the data source type name.
func (r *PortalListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_list"
}

// Schema defines the schema for the data source.
func (r *PortalListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PortalList DataSource",

		Attributes: map[string]schema.Attribute{
			"data": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"application_count": schema.NumberAttribute{
							Computed:    true,
							Description: `Number of applications created in the portal.`,
						},
						"auto_approve_applications": schema.BoolAttribute{
							Computed:    true,
							Description: `Whether the requests from applications to register for products will be automatically approved, or if they will be set to pending until approved by an admin.`,
						},
						"auto_approve_developers": schema.BoolAttribute{
							Computed:    true,
							Description: `Whether the developer account registrations will be automatically approved, or if they will be set to pending until approved by an admin.`,
						},
						"created_at": schema.StringAttribute{
							Computed:    true,
							Description: `An ISO-8601 timestamp representation of entity creation date.`,
						},
						"custom_client_domain": schema.StringAttribute{
							Computed:    true,
							Description: `The custom domain to access a self-hosted customized developer portal client. If this is set, the Konnect-hosted portal client will no longer be available. ` + "`" + `custom_domain` + "`" + ` must be also set for this value to be set. See https://github.com/Kong/konnect-portal for information on how to get started deploying and customizing your own Konnect portal.`,
						},
						"custom_domain": schema.StringAttribute{
							Computed:    true,
							Description: `The custom domain to access the developer portal. A CNAME for the portal's default domain must be able to be set for the custom domain for it to be valid. After setting a valid CNAME, an SSL/TLS certificate will be automatically manged for the custom domain, and traffic will be able to use the custom domain to route to the portal's web client and API.`,
						},
						"default_application_auth_strategy_id": schema.StringAttribute{
							Computed:    true,
							Description: `Default strategy ID applied on applications for the portal`,
						},
						"default_domain": schema.StringAttribute{
							Computed:    true,
							Description: `The domain assigned to the portal by Konnect. This is the default place to access the portal and its API if not using a ` + "`" + `custom_domain` + "``" + `.`,
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: `The description of the portal.`,
						},
						"developer_count": schema.NumberAttribute{
							Computed:    true,
							Description: `Number of developers using the portal.`,
						},
						"display_name": schema.StringAttribute{
							Computed:    true,
							Description: `The display name of the portal. This value will be the portal's ` + "`" + `name` + "`" + ` in Portal API.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `Contains a unique identifier used for this resource.`,
						},
						"is_public": schema.BoolAttribute{
							Computed:    true,
							Description: `Whether the portal catalog can be accessed publicly without any developer authentication. Developer accounts and applications cannot be created if the portal is public.`,
						},
						"labels": schema.MapAttribute{
							Computed:    true,
							ElementType: types.StringType,
							MarkdownDescription: `description: A maximum of 50 user-defined labels are allowed on this resource.` + "\n" +
								`Keys must not start with kong, konnect, insomnia, mesh, kic or _, which are reserved for Kong.` + "\n" +
								`Keys are case-sensitive.` + "\n" +
								``,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `The name of the portal, used to distinguish it from other portals. Name must be unique.`,
						},
						"published_product_count": schema.NumberAttribute{
							Computed:    true,
							Description: `Number of api products published to the portal`,
						},
						"rbac_enabled": schema.BoolAttribute{
							Computed:    true,
							Description: `Whether the portal resources are protected by Role Based Access Control (RBAC). If enabled, developers view or register for products until unless assigned to teams with access to view and consume specific products.`,
						},
						"updated_at": schema.StringAttribute{
							Computed:    true,
							Description: `An ISO-8601 timestamp representation of entity update date.`,
						},
					},
				},
			},
			"meta": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"page": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"number": schema.NumberAttribute{
								Computed: true,
							},
							"size": schema.NumberAttribute{
								Computed: true,
							},
							"total": schema.NumberAttribute{
								Computed: true,
							},
						},
						Description: `Contains pagination query parameters and the total number of objects returned.`,
					},
				},
				Description: `returns the pagination information`,
			},
			"page_number": schema.Int64Attribute{
				Optional:    true,
				Description: `Determines which page of the entities to retrieve.`,
			},
			"page_size": schema.Int64Attribute{
				Optional:    true,
				Description: `The maximum number of items to include per page. The last page of a collection may include fewer items.`,
			},
			"sort": schema.StringAttribute{
				Optional: true,
				MarkdownDescription: `Sorts a collection of portals. Supported sort attributes are:` + "\n" +
					`  - name` + "\n" +
					`  - description` + "\n" +
					`  - is_public` + "\n" +
					`  - rbac_enabled` + "\n" +
					`  - auto_approve_applications` + "\n" +
					`  - auto_approve_developers` + "\n" +
					`  - default_domain` + "\n" +
					`  - custom_domain` + "\n" +
					`  - custom_client_domain` + "\n" +
					`  - created_at` + "\n" +
					`  - updated_at` + "\n" +
					``,
			},
		},
	}
}

func (r *PortalListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PortalListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PortalListDataSourceModel
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

	pageSize := new(int64)
	if !data.PageSize.IsUnknown() && !data.PageSize.IsNull() {
		*pageSize = data.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	pageNumber := new(int64)
	if !data.PageNumber.IsUnknown() && !data.PageNumber.IsNull() {
		*pageNumber = data.PageNumber.ValueInt64()
	} else {
		pageNumber = nil
	}
	sort := new(string)
	if !data.Sort.IsUnknown() && !data.Sort.IsNull() {
		*sort = data.Sort.ValueString()
	} else {
		sort = nil
	}
	request := operations.ListPortalsRequest{
		PageSize:   pageSize,
		PageNumber: pageNumber,
		Sort:       sort,
	}
	res, err := r.client.Portals.ListPortals(ctx, request)
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
	if !(res.ListPortalsResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedListPortalsResponse(res.ListPortalsResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
