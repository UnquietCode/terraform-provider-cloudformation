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

func ResourceAppMeshVirtualRouter() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppMeshVirtualRouterCreate,
		Read:   resourceAppMeshVirtualRouterRead,
		Update: resourceAppMeshVirtualRouterUpdate,
		Delete: resourceAppMeshVirtualRouterDelete,

		Schema: map[string]*schema.Schema{
			"mesh_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"virtual_router_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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
				Required: false,
			},
		},
	}
}

func resourceAppMeshVirtualRouterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppMesh::VirtualRouter", data, meta)
}

func resourceAppMeshVirtualRouterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppMesh::VirtualRouter", data, meta)
}

func resourceAppMeshVirtualRouterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppMesh::VirtualRouter", data, meta)
}

func resourceAppMeshVirtualRouterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppMesh::VirtualRouter", data, meta)
}