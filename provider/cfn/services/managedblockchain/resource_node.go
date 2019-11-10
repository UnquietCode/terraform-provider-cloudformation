// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceNodeCreate,
		Read:   resourceNodeRead,
		Update: resourceNodeUpdate,
		Delete: resourceNodeDelete,

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
		},
	}
}

func resourceNodeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ManagedBlockchain::Node", data, meta)
}

func resourceNodeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ManagedBlockchain::Node", data, meta)
}

func resourceNodeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ManagedBlockchain::Node", data, meta)
}

func resourceNodeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ManagedBlockchain::Node", data, meta)
}