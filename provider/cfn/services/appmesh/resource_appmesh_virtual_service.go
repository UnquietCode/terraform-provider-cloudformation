// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html

package appmesh

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppMeshVirtualService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppMeshVirtualServiceCreate,
		Read:   resourceAppMeshVirtualServiceRead,
		Update: resourceAppMeshVirtualServiceUpdate,
		Delete: resourceAppMeshVirtualServiceDelete,

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
			"virtual_service_name": {
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
				Elem: propertyVirtualServiceVirtualServiceSpec(),
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

func resourceAppMeshVirtualServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppMesh::VirtualService", ResourceAppMeshVirtualService(), data, meta)
}

func resourceAppMeshVirtualServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppMesh::VirtualService", ResourceAppMeshVirtualService(), data, meta)
}

func resourceAppMeshVirtualServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppMesh::VirtualService", ResourceAppMeshVirtualService(), data, meta)
}

func resourceAppMeshVirtualServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppMesh::VirtualService", data, meta)
}