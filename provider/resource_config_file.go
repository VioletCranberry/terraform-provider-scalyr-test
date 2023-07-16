package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/VioletCranberry/terraform-provider-scalyr-test/client"
)

func resourceConfigFile() *schema.Resource {
	return &schema.Resource{
		Description: "Provides JSON resource to create and manage Scalyr config files",
		Schema: map[string]*schema.Schema{
			"file_path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The path to the file at https://app.scalyr.com/files",
			},
			"content": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsJSON,
				Description:  "Json file content",
			},
		},
		CreateContext: resourceConfigFileCreate,
		ReadContext:   resourceConfigFileRead,
		UpdateContext: resourceConfigFileUpdate,
		DeleteContext: resourceConfigFileDelete,
	}
}

func setupConfigFileResource(d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.ApiClient)

	fileApiPath := d.Get("file_path").(string)
	fileContent := d.Get("content").(string)
	r, err := c.CreateConfigFile(
		fileApiPath,
		fileContent)

	if err != nil {
		return diag.FromErr(err)
	}
	if r.Status != "success" {
		return diag.FromErr(fmt.Errorf("invalid API response! message: %s", r.Message))
	}
	d.SetId(stringHash(fileContent))
	return diags
}

func resourceConfigFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return setupConfigFileResource(d, m)
}

func resourceConfigFileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return setupConfigFileResource(d, m)
}

func resourceConfigFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceConfigFileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.ApiClient)

	fileApiPath := d.Get("file_path").(string)
	// do not delete for now - simply overwrite the file
	r, err := c.CreateConfigFile(
		fileApiPath,
		"{}")

	if err != nil {
		return diag.FromErr(err)
	}
	if r.Status != "success" {
		return diag.FromErr(fmt.Errorf("invalid API response! message: %s", r.Message))
	}
	return diags
}
