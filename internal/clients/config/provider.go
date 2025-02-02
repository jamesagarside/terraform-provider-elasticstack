package config

import "github.com/hashicorp/terraform-plugin-framework/types"

type ProviderConfiguration struct {
	Elasticsearch []ElasticsearchConnection `tfsdk:"elasticsearch"`
	Kibana        []KibanaConnection        `tfsdk:"kibana"`
	Fleet         []FleetConnection         `tfsdk:"fleet"`
}

type ElasticsearchConnection struct {
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
	APIKey    types.String `tfsdk:"api_key"`
	Endpoints types.List   `tfsdk:"endpoints"`
	Insecure  types.Bool   `tfsdk:"insecure"`
	CAFile    types.String `tfsdk:"ca_file"`
	CAData    types.String `tfsdk:"ca_data"`
	CertFile  types.String `tfsdk:"cert_file"`
	KeyFile   types.String `tfsdk:"key_file"`
	CertData  types.String `tfsdk:"cert_data"`
	KeyData   types.String `tfsdk:"key_data"`
}

type KibanaConnection struct {
	Username  types.String `tfsdk:"username"`
	Password  types.String `tfsdk:"password"`
	Endpoints types.List   `tfsdk:"endpoints"`
	Insecure  types.Bool   `tfsdk:"insecure"`
}

type FleetConnection struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	APIKey   types.String `tfsdk:"api_key"`
	Endpoint types.String `tfsdk:"endpoint"`
	Insecure types.Bool   `tfsdk:"insecure"`
	CACerts  types.List   `tfsdk:"ca_certs"`
}
