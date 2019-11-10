// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticsearch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDomainEBSOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ebs_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"iops": {
				Type: schema.TypeInt,
				Required: false,
			},
			"volume_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"volume_type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}