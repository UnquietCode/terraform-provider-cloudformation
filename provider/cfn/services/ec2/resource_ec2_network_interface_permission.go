// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfacepermission.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2NetworkInterfacePermissionType string = "AWS::EC2::NetworkInterfacePermission"

func ResourceEC2NetworkInterfacePermission() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2NetworkInterfacePermissionExists,
		Read: resourceEC2NetworkInterfacePermissionRead,
		Create: resourceEC2NetworkInterfacePermissionCreate,
		Update: resourceEC2NetworkInterfacePermissionUpdate,
		Delete: resourceEC2NetworkInterfacePermissionDelete,
		CustomizeDiff: resourceEC2NetworkInterfacePermissionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"aws_account_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"permission": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2NetworkInterfacePermissionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2NetworkInterfacePermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NetworkInterfacePermissionType, ResourceEC2NetworkInterfacePermission(), data, meta)
}

func resourceEC2NetworkInterfacePermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2NetworkInterfacePermissionType, ResourceEC2NetworkInterfacePermission(), data, meta)
}

func resourceEC2NetworkInterfacePermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2NetworkInterfacePermissionType, ResourceEC2NetworkInterfacePermission(), data, meta)
}

func resourceEC2NetworkInterfacePermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2NetworkInterfacePermissionType, data, meta)
}

func resourceEC2NetworkInterfacePermissionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2NetworkInterfacePermissionType, data, meta)
}
