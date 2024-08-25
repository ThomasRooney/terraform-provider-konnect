// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PortalResource{}
var _ resource.ResourceWithImportState = &PortalResource{}

func NewPortalResource() resource.Resource {
	return &PortalResource{}
}

// PortalResource defines the resource implementation.
type PortalResource struct {
	client *sdk.Konnect
}

// PortalResourceModel describes the resource data model.
type PortalResourceModel struct {
	ApplicationCount                 types.Number            `tfsdk:"application_count"`
	AutoApproveApplications          types.Bool              `tfsdk:"auto_approve_applications"`
	AutoApproveDevelopers            types.Bool              `tfsdk:"auto_approve_developers"`
	CreatedAt                        types.String            `tfsdk:"created_at"`
	CustomClientDomain               types.String            `tfsdk:"custom_client_domain"`
	CustomDomain                     types.String            `tfsdk:"custom_domain"`
	DefaultApplicationAuthStrategyID types.String            `tfsdk:"default_application_auth_strategy_id"`
	DefaultDomain                    types.String            `tfsdk:"default_domain"`
	Description                      types.String            `tfsdk:"description"`
	DeveloperCount                   types.Number            `tfsdk:"developer_count"`
	DisplayName                      types.String            `tfsdk:"display_name"`
	Force                            types.String            `tfsdk:"force"`
	ID                               types.String            `tfsdk:"id"`
	IsPublic                         types.Bool              `tfsdk:"is_public"`
	Labels                           map[string]types.String `tfsdk:"labels"`
	Name                             types.String            `tfsdk:"name"`
	PublishedProductCount            types.Number            `tfsdk:"published_product_count"`
	RbacEnabled                      types.Bool              `tfsdk:"rbac_enabled"`
	UpdatedAt                        types.String            `tfsdk:"updated_at"`
}

func (r *PortalResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal"
}

func (r *PortalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Portal Resource",
		Attributes: map[string]schema.Attribute{
			"application_count": schema.NumberAttribute{
				Computed:    true,
				Description: `Number of applications created in the portal.`,
			},
			"auto_approve_applications": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the requests from applications to register for products will be automatically approved, or if they will be set to pending until approved by an admin.`,
			},
			"auto_approve_developers": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the developer account registrations will be automatically approved, or if they will be set to pending until approved by an admin.`,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"custom_client_domain": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The custom domain to access a self-hosted customized developer portal client. If this is set, the Konnect-hosted portal will no longer be available.  ` + "`" + `custom_domain` + "`" + ` must be also set for this value to be set. See https://github.com/Kong/konnect-portal for information on how to get started deploying and customizing your own Konnect portal.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(1024),
				},
			},
			"custom_domain": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The custom domain to access the developer portal. A CNAME for the portal's default domain must be able to be set for the custom domain for it to be valid. After setting a valid CNAME, an SSL/TLS certificate will be automatically manged for the custom domain, and traffic will be able to use the custom domain to route to the portal's web client and API.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(1024),
				},
			},
			"default_application_auth_strategy_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Default strategy ID applied on applications for the portal`,
			},
			"default_domain": schema.StringAttribute{
				Computed:    true,
				Description: `The domain assigned to the portal by Konnect. This is the default place to access the portal and its API if not using a ` + "`" + `custom_domain` + "``" + `.`,
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The description of the portal.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(512),
				},
			},
			"developer_count": schema.NumberAttribute{
				Computed:    true,
				Description: `Number of developers using the portal.`,
			},
			"display_name": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `The display name of the portal. This value will be the portal's ` + "`" + `name` + "`" + ` in Portal API.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthBetween(1, 255),
				},
			},
			"force": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString("false"),
				Description: `If true, delete specified portal and all related entities, even if there are developers registered to portal or if there are portal product versions with application registration enabled. If false, do not allow deletion if there are developers registered to portal or if there are portal product versions with application registration enabled. must be one of ["true", "false"]; Default: "false"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"true",
						"false",
					),
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `Contains a unique identifier used for this resource.`,
			},
			"is_public": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the portal catalog can be accessed publicly without any developer authentication. Developer accounts and applications cannot be created if the portal is public.`,
			},
			"labels": schema.MapAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				MarkdownDescription: `Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types. ` + "\n" +
					`` + "\n" +
					`Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".` + "\n" +
					``,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The name of the portal, used to distinguish it from other portals. Name must be unique.`,
				Validators: []validator.String{
					stringvalidator.UTF8LengthBetween(1, 255),
				},
			},
			"published_product_count": schema.NumberAttribute{
				Computed:    true,
				Description: `Number of api products published to the portal`,
			},
			"rbac_enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the portal resources are protected by Role Based Access Control (RBAC). If enabled, developers view or register for products until unless assigned to teams with access to view and consume specific products.`,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An ISO-8601 timestamp representation of entity update date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *PortalResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PortalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PortalResourceModel
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

	request := *data.ToSharedCreatePortalRequest()
	res, err := r.client.Portals.CreatePortal(ctx, request)
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
	if !(res.CreatePortalResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCreatePortalResponse(res.CreatePortalResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PortalResourceModel
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PortalResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var portalID string
	portalID = data.ID.ValueString()

	updatePortalRequest := *data.ToSharedUpdatePortalRequest()
	request := operations.UpdatePortalRequest{
		PortalID:            portalID,
		UpdatePortalRequest: updatePortalRequest,
	}
	res, err := r.client.Portals.UpdatePortal(ctx, request)
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
	if !(res.UpdatePortalResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUpdatePortalResponse(res.UpdatePortalResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PortalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PortalResourceModel
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

	var portalID string
	portalID = data.ID.ValueString()

	force := new(operations.Force)
	if !data.Force.IsUnknown() && !data.Force.IsNull() {
		*force = operations.Force(data.Force.ValueString())
	} else {
		force = nil
	}
	request := operations.DeletePortalRequest{
		PortalID: portalID,
		Force:    force,
	}
	res, err := r.client.Portals.DeletePortal(ctx, request)
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

func (r *PortalResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource portal.")
}
