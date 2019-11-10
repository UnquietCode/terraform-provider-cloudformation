// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceAppMeshMesh() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppMeshMeshCreate,
		Read:   resourceAppMeshMeshRead,
		Update: resourceAppMeshMeshUpdate,
		Delete: resourceAppMeshMeshDelete,

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
			"spec": {
				Type: schema.TypeList,
				Elem: propertyMeshMeshSpec(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceAppMeshMeshCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppMesh::Mesh", ResourceAppMeshMesh(), data, meta)
}

func resourceAppMeshMeshRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppMesh::Mesh", ResourceAppMeshMesh(), data, meta)
}

func resourceAppMeshMeshUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppMesh::Mesh", ResourceAppMeshMesh(), data, meta)
}

func resourceAppMeshMeshDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppMesh::Mesh", data, meta)
}