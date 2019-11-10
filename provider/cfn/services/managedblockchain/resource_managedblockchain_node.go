// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceManagedBlockchainNodeCreate,
		Read:   resourceManagedBlockchainNodeRead,
		Update: resourceManagedBlockchainNodeUpdate,
		Delete: resourceManagedBlockchainNodeDelete,

		Schema: map[string]*schema.Schema{
			"member_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"node_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
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
		},
	}
}

func resourceManagedBlockchainNodeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ManagedBlockchain::Node", data, meta)
}

func resourceManagedBlockchainNodeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ManagedBlockchain::Node", data, meta)
}

func resourceManagedBlockchainNodeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ManagedBlockchain::Node", data, meta)
}

func resourceManagedBlockchainNodeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ManagedBlockchain::Node", data, meta)
}