package fram

import (
	"context"

	"github.com/darkedges/fram-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBaseURLSource() *schema.Resource {
	return &schema.Resource{
		Description: "Returns details about the [Base URL Source Service](https://backstage.forgerock.com/docs/am/6.5/oidc1-guide/index.html#configure-base-url-source)",
		ReadContext: dataSourceBaseURLSourceRead,
		Schema: map[string]*schema.Schema{
			"source": {
				Type:     schema.TypeString,
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
			"context_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the context path for the base URL.",
			},
			"fixed_value": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If Fixed value is selected as the Base URL source, the base URL in the Fixed value base URL field.",
			},
			"extension_class_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If Extension class is selected as the Base URL source, the Extension class name field.",
			},
		},
	}
}

func dataSourceBaseURLSourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*fram.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bus, err := c.GetBaseURLSource()
	if err != nil {
		if err.Error() == "status: 404, body: {\"code\":404,\"reason\":\"Not Found\",\"message\":\"Not Found\"}" {
			return diags
		} else {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("context_path", bus.Contextpath); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("extension_class_name", bus.ExtensionClassName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fixed_value", bus.FixedValue); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("source", bus.Source); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("test.json")
	return diags
}
