package regru

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const defaultApiEndpoint = "https://api.reg.ru/api/regru2/"

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("REGRU_API_USERNAME", nil),
				Description: "API username for reg.ru",
			},
			"api_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("REGRU_API_PASSWORD", nil),
				Description: "API password for reg.ru",
			},
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultApiEndpoint,
				Description: "reg.ru API endpoint",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"regru_dns_record": resourceRegruDNSRecord(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("api_username").(string)
	password := d.Get("api_password").(string)
	endpoint := d.Get("api_endpoint").(string)

	if (username != "") && (password != "") {
		c := NewClient(username, password, endpoint)

		return c, nil
	}

	return nil, errors.New("empty username and password not allowed")
}
