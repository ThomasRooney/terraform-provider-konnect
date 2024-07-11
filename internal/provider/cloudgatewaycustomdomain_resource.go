// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-konnect/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-konnect/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CloudGatewayCustomDomainResource{}
var _ resource.ResourceWithImportState = &CloudGatewayCustomDomainResource{}

func NewCloudGatewayCustomDomainResource() resource.Resource {
	return &CloudGatewayCustomDomainResource{}
}

// CloudGatewayCustomDomainResource defines the resource implementation.
type CloudGatewayCustomDomainResource struct {
	client *sdk.Konnect
}

// CloudGatewayCustomDomainResourceModel describes the resource data model.
type CloudGatewayCustomDomainResourceModel struct {
	CertificateID   types.String                      `tfsdk:"certificate_id"`
	ControlPlaneGeo types.String                      `tfsdk:"control_plane_geo"`
	ControlPlaneID  types.String                      `tfsdk:"control_plane_id"`
	CreatedAt       types.String                      `tfsdk:"created_at"`
	Domain          types.String                      `tfsdk:"domain"`
	EntityVersion   types.Int64                       `tfsdk:"entity_version"`
	ID              types.String                      `tfsdk:"id"`
	SniID           types.String                      `tfsdk:"sni_id"`
	State           types.String                      `tfsdk:"state"`
	StateMetadata   tfTypes.CustomDomainStateMetadata `tfsdk:"state_metadata"`
	UpdatedAt       types.String                      `tfsdk:"updated_at"`
}

func (r *CloudGatewayCustomDomainResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cloud_gateway_custom_domain"
}

func (r *CloudGatewayCustomDomainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CloudGatewayCustomDomain Resource",
		Attributes: map[string]schema.Attribute{
			"certificate_id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `Certificate ID for the certificate representing this domain and stored on data-planes for this` + "\n" +
					`control-plane. Can be retrieved via the control-planes API for this custom domain's control-plane.` + "\n" +
					``,
			},
			"control_plane_geo": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Set of control-plane geos supported for deploying cloud-gateways configurations. Requires replacement if changed. ; must be one of ["us", "eu", "au"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"us",
						"eu",
						"au",
					),
				},
			},
			"control_plane_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Requires replacement if changed. `,
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of custom domain creation date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"domain": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Domain name of the custom domain. Requires replacement if changed. `,
			},
			"entity_version": schema.Int64Attribute{
				Computed: true,
				MarkdownDescription: `Monotonically-increasing version count of the custom domain, to indicate the order of updates to the custom` + "\n" +
					`domain.` + "\n" +
					``,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"sni_id": schema.StringAttribute{
				Computed: true,
				MarkdownDescription: `Server Name Indication ID for this domain and stored on data-planes for this control-plane. Can be retrieved` + "\n" +
					`via the control-planes API for this custom domain's control-plane.` + "\n" +
					``,
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: `State of the custom domain. must be one of ["created", "initializing", "ready", "terminating", "terminated", "error"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"created",
						"initializing",
						"ready",
						"terminating",
						"terminated",
						"error",
					),
				},
			},
			"state_metadata": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"reason": schema.StringAttribute{
						Computed: true,
						MarkdownDescription: `Reason why the custom domain may be in an erroneous state, reported from backing infrastructure.` + "\n" +
							``,
					},
					"reported_status": schema.StringAttribute{
						Computed:    true,
						Description: `Reported status of the custom domain from backing infrastructure.`,
					},
				},
				MarkdownDescription: `Metadata describing the backing state of the custom domain and why it may be in an erroneous state.` + "\n" +
					``,
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: `An RFC-3339 timestamp representation of custom domain update date.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
		},
	}
}

func (r *CloudGatewayCustomDomainResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CloudGatewayCustomDomainResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CloudGatewayCustomDomainResourceModel
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

	request := *data.ToSharedCreateCustomDomainRequest()
	res, err := r.client.CustomDomains.CreateCustomDomains(ctx, request)
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
	if !(res.CustomDomain != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCustomDomain(res.CustomDomain)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayCustomDomainResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CloudGatewayCustomDomainResourceModel
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

	customDomainID := data.ID.ValueString()
	request := operations.GetCustomDomainRequest{
		CustomDomainID: customDomainID,
	}
	res, err := r.client.CustomDomains.GetCustomDomain(ctx, request)
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
	if !(res.CustomDomain != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCustomDomain(res.CustomDomain)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayCustomDomainResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CloudGatewayCustomDomainResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CloudGatewayCustomDomainResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CloudGatewayCustomDomainResourceModel
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

	customDomainID := data.ID.ValueString()
	request := operations.DeleteCustomDomainRequest{
		CustomDomainID: customDomainID,
	}
	res, err := r.client.CustomDomains.DeleteCustomDomain(ctx, request)
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

func (r *CloudGatewayCustomDomainResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
