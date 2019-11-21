// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-mesh.html

package appmesh

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appMeshMeshType string = "AWS::AppMesh::Mesh"

func ResourceAppMeshMesh() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceAppMeshMeshRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appMeshMeshType, ResourceAppMeshMesh(), data, meta)
}

func resourceAppMeshMeshCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appMeshMeshType, ResourceAppMeshMesh(), data, meta)
}

func resourceAppMeshMeshUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appMeshMeshType, ResourceAppMeshMesh(), data, meta)
}

func resourceAppMeshMeshDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appMeshMeshType, data, meta)
}

func resourceAppMeshMeshCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appMeshMeshType, data, meta)
}
