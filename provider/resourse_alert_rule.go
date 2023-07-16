package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAlertRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Description of the condition that triggers the alert",
			},
			"alert_address": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "One or more email addresses to notify",
			},
			"grace_period_minutes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Duration a triggered event must last before an alert is generated",
			},
			"renotify_period_minutes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Interval, in minutes, at which we send repeat alerts for a continuing alert",
			},
			"resolution_delay_minutes": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Duration an alert must no longer be true to be considered resolved",
			},
			"trigger": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The condition that triggers the alert",
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		CreateContext: resourceAlertRuleCreate,
		ReadContext:   resourceAlertRuleRead,
		UpdateContext: resourceAlertRuleUpdate,
		DeleteContext: resourceAlertRuleDelete,
	}
}

type ResourceAlertRule struct {
	Description            string `json:"description,omitempty"`
	AlertAddress           string `json:"alertAddress,omitempty"`
	GracePeriodMinutes     int    `json:"gracePeriodMinutes,omitempty"`
	RenotifyPeriodMinutes  int    `json:"renotifyPeriodMinutes,omitempty"`
	ResolutionDelayMinutes int    `json:"resolutionDelayMinutes,omitempty"`
	Trigger                string `json:"trigger"`
}

func NewResourceAlertRule(d *schema.ResourceData) *ResourceAlertRule {
	return &ResourceAlertRule{
		AlertAddress:           schemaListToStr(d, "alert_address"),
		Description:            d.Get("description").(string),
		GracePeriodMinutes:     d.Get("grace_period_minutes").(int),
		RenotifyPeriodMinutes:  d.Get("renotify_period_minutes").(int),
		ResolutionDelayMinutes: d.Get("resolution_delay_minutes").(int),
		Trigger:                d.Get("trigger").(string),
	}
}

func setupAlertRuleResource(d *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics
	var res_json string
	resource := NewResourceAlertRule(d)

	trigger := d.Get("trigger").(string)
	res_json, diags = jsonEncode(resource)

	d.Set("json", res_json)
	d.SetId(stringHash(trigger))
	return diags
}

func resourceAlertRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return setupAlertRuleResource(d)
}

func resourceAlertRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return setupAlertRuleResource(d)
}

func resourceAlertRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// empty
	return diags
}

func resourceAlertRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
