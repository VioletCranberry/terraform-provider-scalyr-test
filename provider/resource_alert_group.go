package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAlertGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Description of the alert group",
			},
			"alert_address": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "One or more email addresses to notify",
			},
			"alert_rule": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Description of the condition that triggers the alert",
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
					},
				},
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		CreateContext: resourceAlertGroupCreate,
		ReadContext:   resourceAlertGroupRead,
		UpdateContext: resourceAlertGroupUpdate,
		DeleteContext: resourceAlertGroupDelete,
	}
}

type ResourceAlertGroup struct {
	Alerts []ResourceAlertRule `json:"alerts"`
}

func NewResourceAlertGroup(d *schema.ResourceData) *ResourceAlertGroup {
	alert_address := schemaListToStr(d, "alert_address")
	resource_alert_group := &ResourceAlertGroup{}

	rules := d.Get("alert_rule").(*schema.Set).List()
	for _, alert := range rules {
		rule, ok := alert.(map[string]interface{})
		if ok {
			alert_rule := ResourceAlertRule{
				Description:            rule["description"].(string),
				AlertAddress:           alert_address,
				GracePeriodMinutes:     rule["grace_period_minutes"].(int),
				RenotifyPeriodMinutes:  rule["renotify_period_minutes"].(int),
				ResolutionDelayMinutes: rule["resolution_delay_minutes"].(int),
				Trigger:                rule["trigger"].(string),
			}
			resource_alert_group.Alerts = append(resource_alert_group.Alerts, alert_rule)
		}
	}
	return resource_alert_group
}

func setupAlertGroupResource(d *schema.ResourceData) diag.Diagnostics {

	var diags diag.Diagnostics
	var json_b strings.Builder
	var res_json string
	resource := NewResourceAlertGroup(d)
	desc := d.Get("description").(string)

	for _, alert_rule := range resource.Alerts {
		if alert_json, diags := jsonEncode(alert_rule); diags != nil {
			return diags
		} else {
			json_b.WriteString(fmt.Sprintf("%s,", alert_json))
		}
	}
	// trim last comma to avoid invalid Json
	res_json = strings.TrimSuffix(json_b.String(), ",")
	d.Set("json", res_json)
	d.SetId(stringHash(desc))
	return diags

}

func resourceAlertGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return setupAlertGroupResource(d)
}

func resourceAlertGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return setupAlertGroupResource(d)
}

func resourceAlertGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceAlertGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
