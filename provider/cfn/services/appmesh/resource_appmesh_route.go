// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-route.html

package appmesh

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appMeshRouteType string = "AWS::AppMesh::Route"

func ResourceAppMeshRoute() *schema.Resource {
	return &schema.Resource{
		Read: resourceAppMeshRouteRead,
		Create: resourceAppMeshRouteCreate,
		Update: resourceAppMeshRouteUpdate,
		Delete: resourceAppMeshRouteDelete,
		CustomizeDiff: resourceAppMeshRouteCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"mesh_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"virtual_router_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"route_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"spec": {
				Type: schema.TypeSet,
				Elem: propertyRouteRouteSpec(),
				Required: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAppMeshRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appMeshRouteType, ResourceAppMeshRoute(), data, meta)
}

func resourceAppMeshRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appMeshRouteType, ResourceAppMeshRoute(), data, meta)
}

func resourceAppMeshRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appMeshRouteType, ResourceAppMeshRoute(), data, meta)
}

func resourceAppMeshRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appMeshRouteType, data, meta)
}

func resourceAppMeshRouteCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appMeshRouteType, data, meta)
}
