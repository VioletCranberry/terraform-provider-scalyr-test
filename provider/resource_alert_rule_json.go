package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAlertRuleJson() *schema.Resource {
	return &schema.Resource{
		Description: "Provides JSON resource to create and manage Scalyr config files",
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Description of the condition that triggers the alert",
			},
			"json": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsJSON,
				Description:  "Alert json",
			},
		},
		CreateContext: resourceAlertRuleJsonCreate,
		ReadContext:   resourceAlertRuleJsonRead,
		UpdateContext: resourceAlertRuleJsonUpdate,
		DeleteContext: resourceAlertRuleJsonDelete,
	}
}

func setupResourceAlertRule(d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics
	desc := d.Get("description").(string)
	d.SetId(stringHash(desc))
	return diags
}

func resourceAlertRuleJsonCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return setupResourceAlertRule(d)
}

func resourceAlertRuleJsonRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return setupResourceAlertRule(d)
}

func resourceAlertRuleJsonUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// empty
	return diags
}

func resourceAlertRuleJsonDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
