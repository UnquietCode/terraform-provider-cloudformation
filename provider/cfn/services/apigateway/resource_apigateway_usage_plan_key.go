// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayUsagePlanKeyType string = "AWS::ApiGateway::UsagePlanKey"

func ResourceApiGatewayUsagePlanKey() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceApiGatewayUsagePlanKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayUsagePlanKeyType, ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayUsagePlanKeyType, ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayUsagePlanKeyType, ResourceApiGatewayUsagePlanKey(), data, meta)
}

func resourceApiGatewayUsagePlanKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayUsagePlanKeyType, data, meta)
}

func resourceApiGatewayUsagePlanKeyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayUsagePlanKeyType, data, meta)
}
