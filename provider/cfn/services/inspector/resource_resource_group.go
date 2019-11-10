// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package inspector

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceResourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceResourceGroupCreate,
		Read:   resourceResourceGroupRead,
		Update: resourceResourceGroupUpdate,
		Delete: resourceResourceGroupDelete,

		Schema: map[string]*schema.Schema{
			"resource_group_tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceResourceGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Inspector::ResourceGroup", data, meta)
}

func resourceResourceGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Inspector::ResourceGroup", data, meta)
}

func resourceResourceGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Inspector::ResourceGroup", data, meta)
}

func resourceResourceGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Inspector::ResourceGroup", data, meta)
}