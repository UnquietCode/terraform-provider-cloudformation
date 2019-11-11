// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualnode.html

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppMeshVirtualNode() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppMeshVirtualNodeCreate,
		Read:   resourceAppMeshVirtualNodeRead,
		Update: resourceAppMeshVirtualNodeUpdate,
		Delete: resourceAppMeshVirtualNodeDelete,

		Schema: map[string]*schema.Schema{
			"uid": {
				Type: schema.TypeString,
				Computed: true,
			},
			"mesh_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"virtual_node_name": {
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppMeshVirtualNodeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppMesh::VirtualNode", ResourceAppMeshVirtualNode(), data, meta)
}

func resourceAppMeshVirtualNodeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppMesh::VirtualNode", ResourceAppMeshVirtualNode(), data, meta)
}

func resourceAppMeshVirtualNodeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppMesh::VirtualNode", ResourceAppMeshVirtualNode(), data, meta)
}

func resourceAppMeshVirtualNodeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppMesh::VirtualNode", data, meta)
}