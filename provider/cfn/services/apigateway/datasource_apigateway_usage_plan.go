// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayUsagePlanType string = "AWS::ApiGateway::UsagePlan"

func DatasourceApiGatewayUsagePlan() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayUsagePlanRead,
		
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
			"tags": misc.PropertyTags(),
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
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceApiGatewayUsagePlanRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayUsagePlanType, DatasourceApiGatewayUsagePlan(), data, meta)
}
