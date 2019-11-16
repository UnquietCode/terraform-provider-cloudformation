// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appMeshMeshType string = "AWS::AppMesh::Mesh"

var appMeshMeshProperties map[string]string = map[string]string{
	"mesh_name": "MeshName",
	"spec": "Spec",
	"tags": "Tags",
}

func ResourceAppMeshMesh() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppMeshMeshExists,
		Read: resourceAppMeshMeshRead,
		Create: resourceAppMeshMeshCreate,
		Update: resourceAppMeshMeshUpdate,
		Delete: resourceAppMeshMeshDelete,
		CustomizeDiff: resourceAppMeshMeshCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"mesh_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"spec": {
				Type: schema.TypeSet,
				Elem: propertyMeshMeshSpec(),
				Optional: true,
				MaxItems: 1,
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

func resourceAppMeshMeshExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppMeshMeshRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appMeshMeshType, ResourceAppMeshMesh(), data, meta)
}

func resourceAppMeshMeshCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appMeshMeshType, ResourceAppMeshMesh(), data, appMeshMeshProperties, meta)
}

func resourceAppMeshMeshUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appMeshMeshType, ResourceAppMeshMesh(), data, appMeshMeshProperties, meta)
}

func resourceAppMeshMeshDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appMeshMeshType, data, meta)
}

func resourceAppMeshMeshCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appMeshMeshType, data, meta)
}
