// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-managedblockchain-node.html

package managedblockchain

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const managedBlockchainNodeType string = "AWS::ManagedBlockchain::Node"

func ResourceManagedBlockchainNode() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeSet,
				Elem: propertyNodeNodeConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceManagedBlockchainNodeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(managedBlockchainNodeType, ResourceManagedBlockchainNode(), data, meta)
}

func resourceManagedBlockchainNodeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(managedBlockchainNodeType, ResourceManagedBlockchainNode(), data, meta)
}

func resourceManagedBlockchainNodeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(managedBlockchainNodeType, ResourceManagedBlockchainNode(), data, meta)
}

func resourceManagedBlockchainNodeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(managedBlockchainNodeType, data, meta)
}

func resourceManagedBlockchainNodeCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(managedBlockchainNodeType, data, meta)
}
