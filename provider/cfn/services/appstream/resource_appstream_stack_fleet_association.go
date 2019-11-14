// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppStreamStackFleetAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppStreamStackFleetAssociationExists,
		Read: resourceAppStreamStackFleetAssociationRead,
		Create: resourceAppStreamStackFleetAssociationCreate,
		Update: resourceAppStreamStackFleetAssociationUpdate,
		Delete: resourceAppStreamStackFleetAssociationDelete,
		CustomizeDiff: resourceAppStreamStackFleetAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"fleet_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stack_name": {
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

func resourceAppStreamStackFleetAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppStreamStackFleetAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::StackFleetAssociation", ResourceAppStreamStackFleetAssociation(), data, meta)
}

func resourceAppStreamStackFleetAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::StackFleetAssociation", ResourceAppStreamStackFleetAssociation(), data, meta)
}

func resourceAppStreamStackFleetAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::StackFleetAssociation", ResourceAppStreamStackFleetAssociation(), data, meta)
}

func resourceAppStreamStackFleetAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::StackFleetAssociation", data, meta)
}

func resourceAppStreamStackFleetAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

