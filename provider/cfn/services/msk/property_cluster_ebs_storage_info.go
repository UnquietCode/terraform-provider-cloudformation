// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package msk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterEBSStorageInfo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"volume_size": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
		},
	}
}