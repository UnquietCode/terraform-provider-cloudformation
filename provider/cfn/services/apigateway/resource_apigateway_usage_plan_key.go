// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayUsagePlanKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayUsagePlanKeyCreate,
		Read:   resourceApiGatewayUsagePlanKeyRead,
		Delete: resourceApiGatewayUsagePlanKeyDelete,

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
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayUsagePlanKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::UsagePlanKey", ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::UsagePlanKey", ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::UsagePlanKey", ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::UsagePlanKey", data, meta)
}