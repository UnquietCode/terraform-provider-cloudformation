// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTriggerCondition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"crawler_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"state": {
				Type: schema.TypeString,
				Required: false,
			},
			"crawl_state": {
				Type: schema.TypeString,
				Required: false,
			},
			"logical_operator": {
				Type: schema.TypeString,
				Required: false,
			},
			"job_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}