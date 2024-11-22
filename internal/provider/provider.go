// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/darkedges/fram-client-go/fram"
	"github.com/darkedges/terraform-provider-fram/internal/provider/baseurlsource"
	"github.com/darkedges/terraform-provider-fram/internal/provider/serviceaccount"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure ScaffoldingProvider satisfies various provider interfaces.
var _ provider.Provider = &FRAMProvider{}

// FRAMProvider defines the provider implementation.
type FRAMProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// FRAMProviderModel describes the provider data model.
type FRAMProviderModel struct {
	Host       types.String `tfsdk:"host"`
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
	TOTPSecret types.String `tfsdk:"totp_secret"`
	Realm      types.String `tfsdk:"realm"`
}

func (p *FRAMProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ping"
	resp.Version = p.version
}

func (p *FRAMProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				MarkdownDescription: "FRAM Host to connect as, must include the application context i.e `https://internal.example.com/openam`.<BR>The default is `http://localhost:8080/openam`",
				Required:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "FRAM username to connect as.<BR>The default is `amadmin`",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "FRAM Password of username to connect as.<BR>The default is `p4ssw0rd`",
				Required:            true,
			},
			"totp_secret": schema.StringAttribute{
				MarkdownDescription: "If P1AIC and Admin Account provide a TOTP Secret",
				Optional:            true,
			},
			"realm": schema.StringAttribute{
				MarkdownDescription: "FRAM realm to use i.e `/root`.<BR>The default is `/realm`",
				Optional:            true,
			},
		},
	}
}

func (p *FRAMProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data FRAMProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Request", map[string]interface{}{
		"TOTPSecret": data.TOTPSecret.ValueString(),
	})
	client, _ := fram.NewClient(data.Host.ValueStringPointer(), data.Username.ValueStringPointer(), data.Password.ValueStringPointer(), data.Realm.ValueStringPointer(), data.TOTPSecret.ValueStringPointer())
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *FRAMProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		baseurlsource.NewBaseURLSourceResource,
		serviceaccount.NewServiceAccountResource,
	}
}

func (p *FRAMProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		baseurlsource.NewBaseURLSourceDataSource,
		serviceaccount.NewServiceAccountDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &FRAMProvider{
			version: version,
		}
	}
}
