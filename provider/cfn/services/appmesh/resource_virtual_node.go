// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceVirtualNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualNodeCreate,
		Read:   resourceVirtualNodeRead,
		Update: resourceVirtualNodeUpdate,
		Delete: resourceVirtualNodeDelete,

		Schema: map[string]*schema.Schema{
			"mesh_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"spec": {
				Type: schema.TypeList,
				Elem: propertyVirtualNodeVirtualNodeSpec(),
				Required: true,
				MaxItems: 1,
			},
			"virtual_node_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceVirtualNodeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppMesh::VirtualNode", data, meta)
}

func resourceVirtualNodeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppMesh::VirtualNode", data, meta)
}

func resourceVirtualNodeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppMesh::VirtualNode", data, meta)
}

func resourceVirtualNodeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppMesh::VirtualNode", data, meta)
}