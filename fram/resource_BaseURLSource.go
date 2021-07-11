package fram

import (
	"context"
	"log"
	"strings"

	fram "github.com/darkedges/fram-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBaseURLSource() *schema.Resource {
	return &schema.Resource{
		Description: "This helps automate [Configuring the Base URL Source Service](https://backstage.forgerock.com/docs/am/6.5/oidc1-guide/index.html#configure-base-url-source)",

		CreateContext: resourceOrderCreate,
		ReadContext:   resourceOrderRead,
		UpdateContext: resourceOrderUpdate,
		DeleteContext: resourceOrderDelete,
		Schema: map[string]*schema.Schema{
			"source": {
				Type:     schema.TypeString,
				Required: true,
				Description: "Specifies the source of the base URL. Choose from the following:\n\n" +
					"	- Extension class. `EXTENSION_CLASS`\n\n" +
					"		Specifies that the extension class returns a base URL from a provided `HttpServletRequest`. In the Extension class name field, enter org.forgerock.openam.services.baseurl.BaseURLProvider.\n" +
					"	- Fixed value. `FIXED_VALUE`\n\n" +
					"		Specifies that the base URL is retrieved from a specific base URL value. In the Fixed value base URL field, enter the base URL value.\n" +
					"	- Forwarded header. `FORWARDED_HEADER`\n\n" +
					"		Specifies that the base URL is retrieved from a forwarded header field in the HTTP request. The Forwarded HTTP header field is standardized and specified in [RFC7239](https://tools.ietf.org/html/rfc7239).\n" +
					"	- Host/protocol from incoming request. `EXTENSION_CLASS`\n\n" +
					"		Specifies that the hostname, server name, and port are retrieved from the incoming HTTP request.\n" +
					"	- X-Forwarded-* headers. `EXTENSION_CLASS`\n\n" +
					"		Specifies that the base URL is retrieved from non-standard header fields, such as `X-Forwarded-For`, `X-Forwarded-By`, and `X-Forwarded-Proto`.\n",
			},
			"context_path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the context path for the base URL. If provided, the base URL includes the deployment context path appended to the calculated URL. For example, `/openam`.",
			},
			"fixed_value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "If Fixed value is selected as the Base URL source, enter the base URL in the Fixed value base URL field.",
			},
			"extension_class_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "If Extension class is selected as the Base URL source, enter `org.forgerock.openam.services.baseurl.BaseURLProvider` in the Extension class name field.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceOrderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*fram.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	bus := fram.BaseURLSource{}
	bus.Contextpath = d.Get("context_path").(string)
	bus.ExtensionClassName = d.Get("extension_class_name").(string)
	bus.FixedValue = d.Get("fixed_value").(string)
	bus.Source = d.Get("source").(string)

	_, err := c.CreateBaseURLSource(bus)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "status: 201, body:") {
			return diag.FromErr(err)
		}
	}

	resourceOrderRead(ctx, d, m)

	return diags
}

func resourceOrderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	d.SetId(c.BaseURLSourceId())
	return diags
}

func resourceOrderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*fram.Client)
	var changes = false
	log.Println("[INFO] resourceOrderUpdate")
	bus := fram.BaseURLSource{}
	if d.HasChange("context_path") {
		log.Println("[INFO] changes detected: context_path")
		bus.Contextpath = d.Get("context_path").(string)
		changes = true
	}
	if d.HasChange("extension_class_name") {
		log.Println("[INFO] changes detected: extension_class_name")
		bus.ExtensionClassName = d.Get("extension_class_name").(string)
		changes = true
	}
	if d.HasChange("fixed_value") {
		log.Println("[INFO] changes detected: fixed_value")
		bus.FixedValue = d.Get("fixed_value").(string)
		changes = true
	}
	if d.HasChange("source") {
		log.Println("[INFO] changes detected: source")
		bus.Source = d.Get("source").(string)
		changes = true
	}
	if changes {
		_, err := c.CreateBaseURLSource(bus)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceOrderRead(ctx, d, m)
}

func resourceOrderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*fram.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	_, err := c.DeleteBaseURLSource()
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
