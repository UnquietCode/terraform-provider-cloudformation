// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceHealthCheck() *schema.Resource {
	return &schema.Resource{
		Create: resourceHealthCheckCreate,
		Read:   resourceHealthCheckRead,
		Update: resourceHealthCheckUpdate,
		Delete: resourceHealthCheckDelete,

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
				Required: false,
			},
		},
	}
}

func resourceHealthCheckCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53::HealthCheck", data, meta)
}

func resourceHealthCheckRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53::HealthCheck", data, meta)
}

func resourceHealthCheckUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53::HealthCheck", data, meta)
}

func resourceHealthCheckDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53::HealthCheck", data, meta)
}