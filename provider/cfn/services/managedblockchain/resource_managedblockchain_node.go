// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-node.html

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceManagedBlockchainNode() *schema.Resource {
	return &schema.Resource{
		Exists: resourceManagedBlockchainNodeExists,
		Read: resourceManagedBlockchainNodeRead,
		Create: resourceManagedBlockchainNodeCreate,
		Update: resourceManagedBlockchainNodeUpdate,
		Delete: resourceManagedBlockchainNodeDelete,
		CustomizeDiff: resourceManagedBlockchainNodeCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"member_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"network_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"node_configuration": {
				Type: schema.TypeList,
				Elem: propertyNodeNodeConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceManagedBlockchainNodeExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceManagedBlockchainNodeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ManagedBlockchain::Node", ResourceManagedBlockchainNode(), data, meta)
}

func resourceManagedBlockchainNodeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ManagedBlockchain::Node", ResourceManagedBlockchainNode(), data, meta)
}

func resourceManagedBlockchainNodeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ManagedBlockchain::Node", ResourceManagedBlockchainNode(), data, meta)
}

func resourceManagedBlockchainNodeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ManagedBlockchain::Node", data, meta)
}

func resourceManagedBlockchainNodeCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

