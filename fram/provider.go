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
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FRAM_BASEURL", "http://localhost:8080/openam"),
				Description: "FRAM base URL to connect as, must include the application context i.e `https://fram.example.com/openam`.<BR>The default is `http://localhost:8080/openam`",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FRAM_USERNAME", "amadmin"),
				Description: "FRAM username to connect as.<BR>The default is `amadmin`",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FRAM_PASSWORD", "p4ssw0rd"),
				Description: "FRAM password of username to connect as.<BR>The default is `p4ssw0rd`",
			},
			"realm": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FRAM_REALM", "/realm"),
				Description: "FRAM realm to use i.e `/root`.<BR>The default is `/realm`",
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
