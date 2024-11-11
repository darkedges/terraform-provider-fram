package serviceaccount

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
var _ datasource.DataSource = &ServiceAccountDataSource{}

func NewServiceAccountDataSource() datasource.DataSource {
	return &ServiceAccountDataSource{}
}

// ServiceAccountDataSource defines the data source implementation.
type ServiceAccountDataSource struct {
	client *fram.Client
}

type ServiceAccountDataSourceModel struct {
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Scopes        types.List   `tfsdk:"scopes"`
	AccountStatus types.String `tfsdk:"account_status"`
	JWKS          types.String `tfsdk:"jwks"`
}

func (s ServiceAccountDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_p1aic_serviceaccount"
}

func (s ServiceAccountDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PingOne Advanced Identity Cloud Service Account",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description",
			},
			"scopes": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "scopes",
			},
			"account_status": schema.StringAttribute{
				Computed:    true,
				Description: "Account Statyus",
			},
			"jwks": schema.StringAttribute{
				Computed:    true,
				Description: "jwks",
			},
		},
	}
}

func (d *ServiceAccountDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (s ServiceAccountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ServiceAccountDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// For the purposes of this example code, hardcoding a response value to
	// save into the Terraform state.
	data.Name = types.StringValue("Name")
	data.Description = types.StringValue("Description")
	elements := []string{"one", "two"}
	listValue, _ := types.ListValueFrom(ctx, types.StringType, elements)
	data.Scopes = listValue
	data.AccountStatus = types.StringValue("AccountStatus")
	data.JWKS = types.StringValue("JWKS")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
