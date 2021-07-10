package fram

import (
	"context"

	"github.com/darkedges/fram-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FRAM_BASEURL", nil),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FRAM_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FRAM_PASSWORD", nil),
			},
			"realm": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FRAM_REALM", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"fram_baseurlsource": resourceBaseURLSource(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"fram_baseurlsource": dataSourceBaseURLSource(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	baseUrl := d.Get("base_url").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	realm := d.Get("realm").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := fram.NewClient(&baseUrl, &username, &password, &realm)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}
	c, err := fram.NewClient(nil, nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
