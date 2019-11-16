// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayUsagePlanType string = "AWS::ApiGateway::UsagePlan"

func ResourceApiGatewayUsagePlan() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayUsagePlanExists,
		Read: resourceApiGatewayUsagePlanRead,
		Create: resourceApiGatewayUsagePlanCreate,
		Update: resourceApiGatewayUsagePlanUpdate,
		Delete: resourceApiGatewayUsagePlanDelete,
		CustomizeDiff: resourceApiGatewayUsagePlanCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"api_stages": {
				Type: schema.TypeSet,
				Elem: propertyUsagePlanApiStage(),
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"quota": {
				Type: schema.TypeSet,
				Elem: propertyUsagePlanQuotaSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"throttle": {
				Type: schema.TypeSet,
				Elem: propertyUsagePlanThrottleSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"usage_plan_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayUsagePlanExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayUsagePlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayUsagePlanType, ResourceApiGatewayUsagePlan(), data, meta)
}

func resourceApiGatewayUsagePlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayUsagePlanType, ResourceApiGatewayUsagePlan(), data, meta)
}

func resourceApiGatewayUsagePlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayUsagePlanType, ResourceApiGatewayUsagePlan(), data, meta)
}

func resourceApiGatewayUsagePlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayUsagePlanType, data, meta)
}

func resourceApiGatewayUsagePlanCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayUsagePlanType, data, meta)
}
