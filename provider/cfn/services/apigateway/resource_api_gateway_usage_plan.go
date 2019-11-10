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
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"quota": {
				Type: schema.TypeList,
				Elem: propertyUsagePlanQuotaSettings(),
				Required: false,
				MaxItems: 1,
			},
			"throttle": {
				Type: schema.TypeList,
				Elem: propertyUsagePlanThrottleSettings(),
				Required: false,
				MaxItems: 1,
			},
			"usage_plan_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceApiGatewayUsagePlanCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::UsagePlan", data, meta)
}

func resourceApiGatewayUsagePlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::UsagePlan", data, meta)
}

func resourceApiGatewayUsagePlanUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::UsagePlan", data, meta)
}

func resourceApiGatewayUsagePlanDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::UsagePlan", data, meta)
}