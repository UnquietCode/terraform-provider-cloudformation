// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-resourcegroup.html

package inspector

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceInspectorResourceGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceInspectorResourceGroupExists,
		Read: resourceInspectorResourceGroupRead,
		Create: resourceInspectorResourceGroupCreate,
		Update: resourceInspectorResourceGroupUpdate,
		Delete: resourceInspectorResourceGroupDelete,
		CustomizeDiff: resourceInspectorResourceGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"resource_group_tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceInspectorResourceGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceInspectorResourceGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Inspector::ResourceGroup", ResourceInspectorResourceGroup(), data, meta)
}

func resourceInspectorResourceGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Inspector::ResourceGroup", ResourceInspectorResourceGroup(), data, meta)
}

func resourceInspectorResourceGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Inspector::ResourceGroup", ResourceInspectorResourceGroup(), data, meta)
}

func resourceInspectorResourceGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Inspector::ResourceGroup", data, meta)
}

func resourceInspectorResourceGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
