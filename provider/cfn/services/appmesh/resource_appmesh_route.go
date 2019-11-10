// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppMeshRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppMeshRouteCreate,
		Read:   resourceAppMeshRouteRead,
		Update: resourceAppMeshRouteUpdate,
		Delete: resourceAppMeshRouteDelete,

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
			"route_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"spec": {
				Type: schema.TypeList,
				Elem: propertyRouteRouteSpec(),
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

func resourceAppMeshRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppMesh::Route", data, meta)
}

func resourceAppMeshRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppMesh::Route", data, meta)
}

func resourceAppMeshRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppMesh::Route", data, meta)
}

func resourceAppMeshRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppMesh::Route", data, meta)
}