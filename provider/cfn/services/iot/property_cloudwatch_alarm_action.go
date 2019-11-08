// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCloudwatchAlarmAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alarm_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"state_reason": {
				Type: schema.TypeString,
				Required: true,
			},
			"state_value": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}