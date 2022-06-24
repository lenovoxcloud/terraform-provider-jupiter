package jupiter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"auth_token": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("JUPITER_AUTHTOKEN", nil),
				},
				"jupiter_url": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("JUPITER_URL", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"Jupiter_VM":     dataSourceVM(),
				"Jupiter_Volume": dataSourceVolume(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"Jupiter_VM":     resourceVM(),
				"Jupiter_Volume": resourceVolume(),
			},
			ConfigureContextFunc: providerConfigure,
		}

		return p
	}
}

type apiClient struct {
	authToken   string
	jupiter_url string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	jupiter_url := d.Get("jupiter_url").(string)
	authToken := d.Get("auth_token").(string)
	var diags diag.Diagnostics
	conf := apiClient{
		authToken:   authToken,
		jupiter_url: jupiter_url,
	}
	return &conf, diags
}
