// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceUsagePlanKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceUsagePlanKeyCreate,
		Read:   resourceUsagePlanKeyRead,
		Update: resourceUsagePlanKeyUpdate,
		Delete: resourceUsagePlanKeyDelete,

		Schema: map[string]*schema.Schema{
			"key_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"usage_plan_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceUsagePlanKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::UsagePlanKey", data, meta)
}

func resourceUsagePlanKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::UsagePlanKey", data, meta)
}

func resourceUsagePlanKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::UsagePlanKey", data, meta)
}

func resourceUsagePlanKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::UsagePlanKey", data, meta)
}