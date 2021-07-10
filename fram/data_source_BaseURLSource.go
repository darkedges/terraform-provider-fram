package fram

import (
	"context"

	"github.com/darkedges/fram-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBaseURLSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBaseURLSourceRead,
		Schema: map[string]*schema.Schema{
			"source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"context_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fixed_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"extension_class_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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
