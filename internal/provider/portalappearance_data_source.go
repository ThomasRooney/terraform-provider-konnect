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
var _ datasource.DataSource = &PortalAppearanceDataSource{}
var _ datasource.DataSourceWithConfigure = &PortalAppearanceDataSource{}

func NewPortalAppearanceDataSource() datasource.DataSource {
	return &PortalAppearanceDataSource{}
}

// PortalAppearanceDataSource is the data source implementation.
type PortalAppearanceDataSource struct {
	client *sdk.Konnect
}

// PortalAppearanceDataSourceModel describes the data model.
type PortalAppearanceDataSourceModel struct {
	CustomFonts    *tfTypes.NullableAppearanceFonts          `tfsdk:"custom_fonts"`
	CustomTheme    *tfTypes.NullableAppearanceThemeVariables `tfsdk:"custom_theme"`
	Images         *tfTypes.AppearanceImages                 `tfsdk:"images"`
	PortalID       types.String                              `tfsdk:"portal_id"`
	Text           *tfTypes.NullableAppearanceTextVariables  `tfsdk:"text"`
	ThemeName      types.String                              `tfsdk:"theme_name"`
	UseCustomFonts types.Bool                                `tfsdk:"use_custom_fonts"`
}

// Metadata returns the data source type name.
func (r *PortalAppearanceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_portal_appearance"
}

// Schema defines the schema for the data source.
func (r *PortalAppearanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PortalAppearance DataSource",

		Attributes: map[string]schema.Attribute{
			"custom_fonts": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"base": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the font to render in the browser. must be one of ["Roboto", "Inter", "Open Sans", "Lato", "Slabo 27px", "Slabo 13px", "Oswald", "Source Sans Pro", "Montserrat", "Raleway", "PT Sans", "Lora", "Roboto Mono", "Inconsolata", "Source Code Pro", "PT Mono", "Ubuntu Mono", "IBM Plex Mono"]`,
					},
					"code": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the font to render in the browser. must be one of ["Roboto", "Inter", "Open Sans", "Lato", "Slabo 27px", "Slabo 13px", "Oswald", "Source Sans Pro", "Montserrat", "Raleway", "PT Sans", "Lora", "Roboto Mono", "Inconsolata", "Source Code Pro", "PT Mono", "Ubuntu Mono", "IBM Plex Mono"]`,
					},
					"headings": schema.StringAttribute{
						Computed:    true,
						Description: `The name of the font to render in the browser. must be one of ["Roboto", "Inter", "Open Sans", "Lato", "Slabo 27px", "Slabo 13px", "Oswald", "Source Sans Pro", "Montserrat", "Raleway", "PT Sans", "Lora", "Roboto Mono", "Inconsolata", "Source Code Pro", "PT Mono", "Ubuntu Mono", "IBM Plex Mono"]`,
					},
				},
				Description: `Font selections to render text in the portal user interface. Must set use_custom_fonts to true to enable using custom font values.`,
			},
			"custom_theme": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"colors": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"button": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"primary_fill": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"primary_text": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
							"section": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"accent": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"body": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"footer": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"header": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"hero": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"stroke": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"tertiary": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
							"text": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"accent": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"footer": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"header": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"headings": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"hero": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"link": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"primary": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									"secondary": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"description": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
							},
						},
					},
				},
				Description: `Groups of variables for configuring visual details of the portal user interface. Set theme_name to 'custom' to use custom values for theme variables.`,
			},
			"images": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"catalog_cover": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Computed:    true,
								Description: `must be a data URL with base64 image data, e.g., data:image/jpeg;base64,<BASE64_IMAGE_DATA>`,
							},
							"filename": schema.StringAttribute{
								Computed: true,
							},
						},
						Description: `The image data to upload, along with an optional filename. Images must be a data URL with binary image data in base 64 format. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs.`,
					},
					"favicon": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Computed:    true,
								Description: `must be a data URL with base64 image data, e.g., data:image/jpeg;base64,<BASE64_IMAGE_DATA>`,
							},
							"filename": schema.StringAttribute{
								Computed: true,
							},
						},
						Description: `The image data to upload, along with an optional filename. Images must be a data URL with binary image data in base 64 format. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs.`,
					},
					"logo": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Computed:    true,
								Description: `must be a data URL with base64 image data, e.g., data:image/jpeg;base64,<BASE64_IMAGE_DATA>`,
							},
							"filename": schema.StringAttribute{
								Computed: true,
							},
						},
						Description: `The image data to upload, along with an optional filename. Images must be a data URL with binary image data in base 64 format. See https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs.`,
					},
				},
				Description: `A collection of binary image data to customize images in the portal`,
			},
			"portal_id": schema.StringAttribute{
				Required:    true,
				Description: `ID of the portal.`,
			},
			"text": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"catalog": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"primary_header": schema.StringAttribute{
								Computed: true,
							},
							"welcome_message": schema.StringAttribute{
								Computed: true,
							},
						},
					},
				},
				Description: `Values to display for customizable text in the portal user interface`,
			},
			"theme_name": schema.StringAttribute{
				Computed:    true,
				Description: `Select a pre-existing default theme or specify 'custom' to use custom_theme variables. must be one of ["mint_rocket", "dark_mode", "custom"]`,
			},
			"use_custom_fonts": schema.BoolAttribute{
				Computed:    true,
				Description: `If true, fonts in custom_fonts will be used over the theme's default fonts`,
			},
		},
	}
}

func (r *PortalAppearanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PortalAppearanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PortalAppearanceDataSourceModel
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

	portalID := data.PortalID.ValueString()
	request := operations.GetPortalAppearanceRequest{
		PortalID: portalID,
	}
	res, err := r.client.PortalAppearance.GetPortalAppearance(ctx, request)
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
	if !(res.GetPortalAppearanceResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedGetPortalAppearanceResponse(res.GetPortalAppearanceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
