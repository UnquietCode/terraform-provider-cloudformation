// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayUsagePlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayUsagePlanCreate,
		Read:   resourceApiGatewayUsagePlanRead,
		Update: resourceApiGatewayUsagePlanUpdate,
		Delete: resourceApiGatewayUsagePlanDelete,

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
				Type: schema.TypeList,
				Elem: propertyUsagePlanQuotaSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"throttle": {
				Type: schema.TypeList,
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
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayUsagePlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::UsagePlan", ResourceApiGatewayUsagePlan(), data, meta)
}

func resourceApiGatewayUsagePlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::UsagePlan", ResourceApiGatewayUsagePlan(), data, meta)
}

func resourceApiGatewayUsagePlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::UsagePlan", ResourceApiGatewayUsagePlan(), data, meta)
}

func resourceApiGatewayUsagePlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::UsagePlan", data, meta)
}