// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLogConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_watch_logs_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"exclude_verbose_content": {
				Type: schema.TypeBool,
				Required: false,
			},
			"field_log_level": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}