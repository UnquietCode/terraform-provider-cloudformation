// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const appMeshVirtualNodeType string = "AWS::AppMesh::VirtualNode"

func ResourceAppMeshVirtualNode() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppMeshVirtualNodeExists,
		Read: resourceAppMeshVirtualNodeRead,
		Create: resourceAppMeshVirtualNodeCreate,
		Update: resourceAppMeshVirtualNodeUpdate,
		Delete: resourceAppMeshVirtualNodeDelete,
		CustomizeDiff: resourceAppMeshVirtualNodeCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"mesh_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"spec": {
				Type: schema.TypeSet,
				Elem: propertyVirtualNodeVirtualNodeSpec(),
				Required: true,
				MaxItems: 1,
			},
			"virtual_node_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppMeshVirtualNodeExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppMeshVirtualNodeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appMeshVirtualNodeType, ResourceAppMeshVirtualNode(), data, meta)
}

func resourceAppMeshVirtualNodeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appMeshVirtualNodeType, ResourceAppMeshVirtualNode(), data, meta)
}

func resourceAppMeshVirtualNodeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appMeshVirtualNodeType, ResourceAppMeshVirtualNode(), data, meta)
}

func resourceAppMeshVirtualNodeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appMeshVirtualNodeType, data, meta)
}

func resourceAppMeshVirtualNodeCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appMeshVirtualNodeType, data, meta)
}
