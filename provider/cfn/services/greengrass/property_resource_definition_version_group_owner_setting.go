// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyResourceDefinitionVersionGroupOwnerSetting() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_add_group_owner": {
				Type: schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"group_owner": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}