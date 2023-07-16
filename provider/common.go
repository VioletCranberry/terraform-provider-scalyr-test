package provider

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func stringHash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func schemaListToStr(d *schema.ResourceData, schema_field string) string {
	var out strings.Builder

	list := d.Get(schema_field).([]interface{})
	items := make([]string, len(list))
	for i, value := range list {
		items[i] = value.(string)
		out.WriteString(fmt.Sprintf("%s,", items[i]))
	}
	return strings.TrimSuffix(out.String(), ",")
}

func jsonEncode(v interface{}) (string, diag.Diagnostics) {
	var diags diag.Diagnostics
	json_str, err := json.Marshal(v)
	if err != nil {
		return "", diag.FromErr(err)
	}
	
	return string(json_str), diags
}
