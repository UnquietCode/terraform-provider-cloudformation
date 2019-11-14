// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

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
				Type: schema.TypeList,
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
	return plugin.ResourceRead("AWS::Route53::HealthCheck", ResourceRoute53HealthCheck(), data, meta)
}

func resourceRoute53HealthCheckCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53::HealthCheck", ResourceRoute53HealthCheck(), data, meta)
}

func resourceRoute53HealthCheckUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53::HealthCheck", ResourceRoute53HealthCheck(), data, meta)
}

func resourceRoute53HealthCheckDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53::HealthCheck", data, meta)
}

func resourceRoute53HealthCheckCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
