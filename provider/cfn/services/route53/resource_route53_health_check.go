// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53HealthCheckType string = "AWS::Route53::HealthCheck"

var route53HealthCheckProperties map[string]string = map[string]string{
	"health_check_config": "HealthCheckConfig",
	"health_check_tags": "HealthCheckTags",
}

func ResourceRoute53HealthCheck() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoute53HealthCheckExists,
		Read: resourceRoute53HealthCheckRead,
		Create: resourceRoute53HealthCheckCreate,
		Update: resourceRoute53HealthCheckUpdate,
		Delete: resourceRoute53HealthCheckDelete,
		CustomizeDiff: resourceRoute53HealthCheckCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"health_check_config": {
				Type: schema.TypeSet,
				Elem: propertyHealthCheckHealthCheckConfig(),
				Required: true,
				MaxItems: 1,
			},
			"health_check_tags": {
				Type: schema.TypeList,
				Elem: propertyHealthCheckHealthCheckTag(),
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

func resourceRoute53HealthCheckExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoute53HealthCheckRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53HealthCheckType, ResourceRoute53HealthCheck(), data, meta)
}

func resourceRoute53HealthCheckCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(route53HealthCheckType, ResourceRoute53HealthCheck(), data, route53HealthCheckProperties, meta)
}

func resourceRoute53HealthCheckUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(route53HealthCheckType, ResourceRoute53HealthCheck(), data, route53HealthCheckProperties, meta)
}

func resourceRoute53HealthCheckDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(route53HealthCheckType, data, meta)
}

func resourceRoute53HealthCheckCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(route53HealthCheckType, data, meta)
}
