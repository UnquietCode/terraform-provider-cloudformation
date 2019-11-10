// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservicepermissions.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VPCEndpointServicePermissions() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VPCEndpointServicePermissionsCreate,
		Read:   resourceEC2VPCEndpointServicePermissionsRead,
		Update: resourceEC2VPCEndpointServicePermissionsUpdate,
		Delete: resourceEC2VPCEndpointServicePermissionsDelete,

		Schema: map[string]*schema.Schema{
			"allowed_principals": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"service_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2VPCEndpointServicePermissionsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VPCEndpointServicePermissions", ResourceEC2VPCEndpointServicePermissions(), data, meta)
}

func resourceEC2VPCEndpointServicePermissionsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VPCEndpointServicePermissions", ResourceEC2VPCEndpointServicePermissions(), data, meta)
}

func resourceEC2VPCEndpointServicePermissionsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VPCEndpointServicePermissions", ResourceEC2VPCEndpointServicePermissions(), data, meta)
}

func resourceEC2VPCEndpointServicePermissionsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VPCEndpointServicePermissions", data, meta)
}