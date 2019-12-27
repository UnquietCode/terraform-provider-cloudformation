// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html

package route53

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53HealthCheckType string = "AWS::Route53::HealthCheck"

func DatasourceRoute53HealthCheck() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRoute53HealthCheckRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceRoute53HealthCheckRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53HealthCheckType, DatasourceRoute53HealthCheck(), data, meta)
}
