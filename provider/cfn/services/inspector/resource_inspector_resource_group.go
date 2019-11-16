// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const inspectorResourceGroupType string = "AWS::Inspector::ResourceGroup"

var inspectorResourceGroupProperties map[string]string = map[string]string{
	"resource_group_tags": "ResourceGroupTags",
}

func ResourceInspectorResourceGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceInspectorResourceGroupExists,
		Read: resourceInspectorResourceGroupRead,
		Create: resourceInspectorResourceGroupCreate,
		Update: resourceInspectorResourceGroupUpdate,
		Delete: resourceInspectorResourceGroupDelete,
		CustomizeDiff: resourceInspectorResourceGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"resource_group_tags": misc.PropertyTags(),
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
	return plugin.ResourceRead(inspectorResourceGroupType, ResourceInspectorResourceGroup(), data, meta)
}

func resourceInspectorResourceGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(inspectorResourceGroupType, ResourceInspectorResourceGroup(), data, inspectorResourceGroupProperties, meta)
}

func resourceInspectorResourceGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(inspectorResourceGroupType, ResourceInspectorResourceGroup(), data, inspectorResourceGroupProperties, meta)
}

func resourceInspectorResourceGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(inspectorResourceGroupType, data, meta)
}

func resourceInspectorResourceGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(inspectorResourceGroupType, data, meta)
}
