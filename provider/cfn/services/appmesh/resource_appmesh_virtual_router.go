// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualrouter.html

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppMeshVirtualRouter() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppMeshVirtualRouterExists,
		Read: resourceAppMeshVirtualRouterRead,
		Create: resourceAppMeshVirtualRouterCreate,
		Update: resourceAppMeshVirtualRouterUpdate,
		Delete: resourceAppMeshVirtualRouterDelete,
		CustomizeDiff: resourceAppMeshVirtualRouterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"mesh_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"virtual_router_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"spec": {
				Type: schema.TypeList,
				Elem: propertyVirtualRouterVirtualRouterSpec(),
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
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppMeshVirtualRouterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppMeshVirtualRouterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppMesh::VirtualRouter", ResourceAppMeshVirtualRouter(), data, meta)
}

func resourceAppMeshVirtualRouterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppMesh::VirtualRouter", ResourceAppMeshVirtualRouter(), data, meta)
}

func resourceAppMeshVirtualRouterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppMesh::VirtualRouter", ResourceAppMeshVirtualRouter(), data, meta)
}

func resourceAppMeshVirtualRouterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppMesh::VirtualRouter", data, meta)
}

func resourceAppMeshVirtualRouterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
