// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceApiGatewayUsagePlanKeyExists,
		Read: resourceApiGatewayUsagePlanKeyRead,
		Create: resourceApiGatewayUsagePlanKeyCreate,
		Update: resourceApiGatewayUsagePlanKeyUpdate,
		Delete: resourceApiGatewayUsagePlanKeyDelete,
		CustomizeDiff: resourceApiGatewayUsagePlanKeyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"key_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"key_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"usage_plan_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayUsagePlanKeyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayUsagePlanKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::UsagePlanKey", ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::UsagePlanKey", ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::UsagePlanKey", ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::UsagePlanKey", data, meta)
}

func resourceApiGatewayUsagePlanKeyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::ApiGateway::UsagePlanKey", data, meta)
}

