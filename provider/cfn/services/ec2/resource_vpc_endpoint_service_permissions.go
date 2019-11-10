// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceVPCEndpointServicePermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourceVPCEndpointServicePermissionsCreate,
		Read:   resourceVPCEndpointServicePermissionsRead,
		Update: resourceVPCEndpointServicePermissionsUpdate,
		Delete: resourceVPCEndpointServicePermissionsDelete,

		Schema: map[string]*schema.Schema{
			"allowed_principals": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"service_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVPCEndpointServicePermissionsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCEndpointServicePermissions", data, meta)
}

func resourceVPCEndpointServicePermissionsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCEndpointServicePermissions", data, meta)
}

func resourceVPCEndpointServicePermissionsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCEndpointServicePermissions", data, meta)
}

func resourceVPCEndpointServicePermissionsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCEndpointServicePermissions", data, meta)
}