package serviceaccount

import (
	"context"
	"github.com/darkedges/fram-client-go/fram"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ServiceAccountResource{}
var _ resource.ResourceWithImportState = &ServiceAccountResource{}

func NewServiceAccountResource() resource.Resource {
	return &ServiceAccountResource{}
}

// ServiceAccountResource defines the resource implementation.
type ServiceAccountResource struct {
	client *fram.Client
}

func (s ServiceAccountResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_p1aic_serviceaccount"
}

func (s ServiceAccountResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PingOne Advanced Identity Cloud Service Account",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name",
			},
			"description": schema.StringAttribute{
				Required:    true,
				Description: "Description",
			},
			"scopes": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "scopes",
			},
			"account_status": schema.StringAttribute{
				Required:    true,
				Description: "Account Statyus",
			},
			"jwks": schema.StringAttribute{
				Required:    true,
				Description: "jwks",
			},
		},
	}
}

func (s ServiceAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *ServiceAccountResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
