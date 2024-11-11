package serviceaccount

import (
	"context"
	"github.com/darkedges/fram-client-go/fram"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
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

func (s ServiceAccountDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_p1aic_serviceaccount"
}

func (s ServiceAccountDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PingOne Advanced Identity Cloud Service Account",

		Attributes: map[string]schema.Attribute{
			"source": schema.StringAttribute{
				Required:    true,
				Description: "Description",
			},
		},
	}
}

func (s ServiceAccountDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	//TODO implement me
	panic("implement me")
}
