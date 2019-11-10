// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyResourceDefinitionVersionLocalVolumeResourceData() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source_path": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_path": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"group_owner_setting": {
				Type: schema.TypeList,
				Elem: propertyResourceDefinitionVersionGroupOwnerSetting(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}