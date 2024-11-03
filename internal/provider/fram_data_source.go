// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/darkedges/fram-client-go/fram"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &BaseURLSourceDataSource{}

func NewBaseURLSourceDataSource() datasource.DataSource {
	return &BaseURLSourceDataSource{}
}

// BaseURLSourceDataSource defines the data source implementation.
type BaseURLSourceDataSource struct {
	client *fram.Client
}

// BaseURLSourceDataSourceModel describes the data source data model.
type BaseURLSourceDataSourceModel struct {
	Source             types.String `tfsdk:"source"`
	ContextPath        types.String `tfsdk:"context_path"`
	FixedValue         types.String `tfsdk:"fixed_value"`
	ExtensionClassName types.String `tfsdk:"extension_class_name"`
}

func (d *BaseURLSourceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "fram_" + req.ProviderTypeName
}

func (d *BaseURLSourceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example data source",

		Attributes: map[string]schema.Attribute{
			"source": schema.StringAttribute{
				Computed: true,
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
				Computed:    true,
				Description: "Specifies the context path for the base URL. If provided, the base URL includes the deployment context path appended to the calculated URL. For example, `/openam`.",
			},
			"fixed_value": schema.StringAttribute{
				Computed:    true,
				Description: "If Fixed value is selected as the Base URL source, enter the base URL in the Fixed value base URL field.",
			},
			"extension_class_name": schema.StringAttribute{
				Computed:    true,
				Description: "If Extension class is selected as the Base URL source, enter `org.forgerock.openam.services.baseurl.BaseURLProvider` in the Extension class name field.",
			},
		},
	}
}

func (d *BaseURLSourceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*fram.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *BaseURLSourceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data BaseURLSourceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bus, err := d.client.GetBaseURLSource()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
		return
	}

	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.ContextPath = types.StringValue(bus.Contextpath)
	data.Source = types.StringValue(bus.Source)
	data.ExtensionClassName = types.StringValue(bus.ExtensionClassName)
	data.FixedValue = types.StringValue(bus.FixedValue)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

//package provider
//
//import (
//	"context"
//
//	"github.com/darkedges/internal-client-go"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//)
//
//func dataSourceBaseURLSource() *schema.Resource {
//	return &schema.Resource{
//		Description: "Returns details about the [Base URL Source Service](https://backstage.forgerock.com/docs/am/6.5/oidc1-guide/index.html#configure-base-url-source)",
//		ReadContext: dataSourceBaseURLSourceRead,
//		Schema: map[string]*schema.Schema{
//			"source": {
//				Type:     schema.TypeString,
//				Computed: true,
//				Description: "Specifies the source of the base URL. Choose from the following:\n\n" +
//					"	- Extension class. `EXTENSION_CLASS`\n\n" +
//					"		Specifies that the extension class returns a base URL from a provided `HttpServletRequest`. In the Extension class name field, enter org.forgerock.openam.services.baseurl.BaseURLProvider.\n" +
//					"	- Fixed value. `FIXED_VALUE`\n\n" +
//					"		Specifies that the base URL is retrieved from a specific base URL value. In the Fixed value base URL field, enter the base URL value.\n" +
//					"	- Forwarded header. `FORWARDED_HEADER`\n\n" +
//					"		Specifies that the base URL is retrieved from a forwarded header field in the HTTP request. The Forwarded HTTP header field is standardized and specified in [RFC7239](https://tools.ietf.org/html/rfc7239).\n" +
//					"	- Host/protocol from incoming request. `REQUEST_VALUES`\n\n" +
//					"		Specifies that the hostname, server name, and port are retrieved from the incoming HTTP request.\n" +
//					"	- X-Forwarded-* headers. `X_FORWARDED_HEADERS`\n\n" +
//					"		Specifies that the base URL is retrieved from non-standard header fields, such as `X-Forwarded-For`, `X-Forwarded-By`, and `X-Forwarded-Proto`.\n",
//			},
//			"context_path": {
//				Type:        schema.TypeString,
//				Computed:    true,
//				Description: "Specifies the context path for the base URL.",
//			},
//			"fixed_value": {
//				Type:        schema.TypeString,
//				Computed:    true,
//				Description: "If Fixed value is selected as the Base URL source, the base URL in the Fixed value base URL field.",
//			},
//			"extension_class_name": {
//				Type:        schema.TypeString,
//				Computed:    true,
//				Description: "If Extension class is selected as the Base URL source, the Extension class name field.",
//			},
//		},
//	}
//}
//
//func dataSourceBaseURLSourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
//	c := m.(*internal.Client)
//
//	// Warning or errors can be collected in a slice type
//	var diags diag.Diagnostics
//
//	bus, err := c.GetBaseURLSource()
//	if err != nil {
//		if err.Error() == "status: 404, body: {\"code\":404,\"reason\":\"Not Found\",\"message\":\"Not Found\"}" {
//			return diags
//		} else {
//			return diag.FromErr(err)
//		}
//	}
//
//	if err := d.Set("context_path", bus.Contextpath); err != nil {
//		return diag.FromErr(err)
//	}
//	if err := d.Set("extension_class_name", bus.ExtensionClassName); err != nil {
//		return diag.FromErr(err)
//	}
//	if err := d.Set("fixed_value", bus.FixedValue); err != nil {
//		return diag.FromErr(err)
//	}
//	if err := d.Set("source", bus.Source); err != nil {
//		return diag.FromErr(err)
//	}
//	d.SetId("test.json")
//	return diags
//}
