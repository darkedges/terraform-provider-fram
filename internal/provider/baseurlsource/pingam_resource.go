// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package baseurlsource

import (
	"context"
	"fmt"
	"github.com/darkedges/fram-client-go/fram"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"strings"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &BaseURLSourceResource{}
var _ resource.ResourceWithImportState = &BaseURLSourceResource{}

func NewBaseURLSourceResource() resource.Resource {
	return &BaseURLSourceResource{}
}

// BaseURLSourceResource defines the resource implementation.
type BaseURLSourceResource struct {
	client *fram.Client
}

// BaseURLSourceModel describes the resource data model.
type BaseURLSourceModel struct {
	Source             types.String `tfsdk:"source"`
	ContextPath        types.String `tfsdk:"context_path"`
	FixedValue         types.String `tfsdk:"fixed_value"`
	ExtensionClassName types.String `tfsdk:"extension_class_name"`
}

func (r *BaseURLSourceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_am_baseurlsource"
}

func (r *BaseURLSourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PingAM Base URL Source",

		Attributes: map[string]schema.Attribute{
			"source": schema.StringAttribute{
				Required: true,
				Description: "Specifies the source of the base URL. Choose from the following:\n\n" +
					"	- Extension class. `EXTENSION_CLASS`\n\n" +
					"		Specifies that the extension class returns a base URL from a provided `HttpServletRequest`. In the Extension class name field, enter org.forgerock.openam.services.baseurl.BaseURLProvider.\n" +
					"	- Fixed value. `FIXED_VALUE`\n\n" +
					"		Specifies that the base URL is retrieved from a specific base URL value. In the Fixed value base URL field, enter the base URL value.\n" +
					"	- Forwarded header. `FORWARDED_HEADER`\n\n" +
					"		Specifies that the base URL is retrieved from a forwarded header field in the HTTP request. The Forwarded HTTP header field is standardized and specified in [RFC7239](https://tools.ietf.org/html/rfc7239).\n" +
					"	- Host/protocol from incoming request. `REQUEST_VALUES`\n\n" +
					"		Specifies that the hostname, server name, and port are retrieved from the incoming HTTP request.\n" +
					"	- X-Forwarded-* headers. `X_FORWARDED_HEADERS`\n\n" +
					"		Specifies that the base URL is retrieved from non-standard header fields, such as `X-Forwarded-For`, `X-Forwarded-By`, and `X-Forwarded-Proto`.\n",
			},
			"context_path": schema.StringAttribute{
				Required:    true,
				Description: "Specifies the context path for the base URL. If provided, the base URL includes the deployment context path appended to the calculated URL. For example, `/openam`.",
			},
			"fixed_value": schema.StringAttribute{
				Required:    true,
				Description: "If Fixed value is selected as the Base URL source, enter the base URL in the Fixed value base URL field.",
			},
			"extension_class_name": schema.StringAttribute{
				Optional:    true,
				Description: "If Extension class is selected as the Base URL source, enter `org.forgerock.openam.services.baseurl.BaseURLProvider` in the Extension class name field.",
			},
		},
	}
}

func (r *BaseURLSourceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*fram.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *BaseURLSourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BaseURLSourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bus := fram.BaseURLSource{
		Contextpath:        data.ContextPath.ValueString(),
		ExtensionClassName: data.ExtensionClassName.ValueString(),
		FixedValue:         data.FixedValue.ValueString(),
		Source:             data.Source.ValueString(),
	}
	result, err := r.client.CreateBaseURLSource(bus)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "status: 201, body:") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Base URL service, got error: %s", err))
			return
		}
	}

	if result != nil {
		// save into the Terraform state.
		data.ContextPath = types.StringValue(result.Contextpath)
		data.Source = types.StringValue(result.Source)
		data.ExtensionClassName = types.StringValue(result.ExtensionClassName)
		data.FixedValue = types.StringValue(result.FixedValue)
	}
	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BaseURLSourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BaseURLSourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	result, err := r.client.GetBaseURLSource()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Base URL service, got error: %s", err))
		return
	}

	// save into the Terraform state.
	data.ContextPath = types.StringValue(result.Contextpath)
	data.Source = types.StringValue(result.Source)
	data.ExtensionClassName = types.StringValue(result.ExtensionClassName)
	data.FixedValue = types.StringValue(result.FixedValue)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BaseURLSourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BaseURLSourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	bus := fram.BaseURLSource{
		Contextpath:        data.ContextPath.ValueString(),
		ExtensionClassName: data.ExtensionClassName.ValueString(),
		FixedValue:         data.FixedValue.ValueString(),
		Source:             data.Source.ValueString(),
	}
	result, err := r.client.UpdateBaseURLSource(bus)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "status: 201, body:") {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Base URL service, got error: %s", err))
			return
		}
	}
	// save into the Terraform state.
	data.ContextPath = types.StringValue(result.Contextpath)
	data.Source = types.StringValue(result.Source)
	data.ExtensionClassName = types.StringValue(result.ExtensionClassName)
	data.FixedValue = types.StringValue(result.FixedValue)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BaseURLSourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BaseURLSourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.DeleteBaseURLSource()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete example, got error: %s", err))
	}
}

func (r *BaseURLSourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
