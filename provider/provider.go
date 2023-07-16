package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/VioletCranberry/terraform-provider-scalyr-test/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"scalyr_app_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALYR_APP_URL", ""),
				Default:     "https://app.scalyr.com",
			},
			"scalyr_api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SCALYR_API_KEY", ""),
			},
			"client_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_TIMEOUT", ""),
				Default:     60,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"scalyr_alert_rule":      resourceAlertRule(),
			"scalyr_config_file":     resourceConfigFile(),
			"scalyr_alert_group":     resourceAlertGroup(),
			"scalyr_alert_rule_json": resourceAlertRuleJson(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	scalyr_app_url := d.Get("scalyr_app_url").(string)
	scalyr_api_key := d.Get("scalyr_api_key").(string)
	client_timeout := d.Get("client_timeout").(int)

	var diags diag.Diagnostics

	return client.NewApiClient(
			scalyr_app_url,
			scalyr_api_key,
			client_timeout),
		diags
}
