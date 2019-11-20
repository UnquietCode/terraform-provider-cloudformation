// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const appMeshVirtualServiceType string = "AWS::AppMesh::VirtualService"

func ResourceAppMeshVirtualService() *schema.Resource {
	return &schema.Resource{
		Read: resourceAppMeshVirtualServiceRead,
		Create: resourceAppMeshVirtualServiceCreate,
		Update: resourceAppMeshVirtualServiceUpdate,
		Delete: resourceAppMeshVirtualServiceDelete,
		CustomizeDiff: resourceAppMeshVirtualServiceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"mesh_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"virtual_service_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"spec": {
				Type: schema.TypeSet,
				Elem: propertyVirtualServiceVirtualServiceSpec(),
				Required: true,
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

func resourceAppMeshVirtualServiceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appMeshVirtualServiceType, ResourceAppMeshVirtualService(), data, meta)
}

func resourceAppMeshVirtualServiceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appMeshVirtualServiceType, ResourceAppMeshVirtualService(), data, meta)
}

func resourceAppMeshVirtualServiceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appMeshVirtualServiceType, ResourceAppMeshVirtualService(), data, meta)
}

func resourceAppMeshVirtualServiceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appMeshVirtualServiceType, data, meta)
}

func resourceAppMeshVirtualServiceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appMeshVirtualServiceType, data, meta)
}
