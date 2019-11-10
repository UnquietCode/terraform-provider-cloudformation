// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppStreamStackFleetAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppStreamStackFleetAssociationCreate,
		Read:   resourceAppStreamStackFleetAssociationRead,
		Update: resourceAppStreamStackFleetAssociationUpdate,
		Delete: resourceAppStreamStackFleetAssociationDelete,

		Schema: map[string]*schema.Schema{
			"fleet_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stack_name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAppStreamStackFleetAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::StackFleetAssociation", data, meta)
}

func resourceAppStreamStackFleetAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::StackFleetAssociation", data, meta)
}

func resourceAppStreamStackFleetAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::StackFleetAssociation", data, meta)
}

func resourceAppStreamStackFleetAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::StackFleetAssociation", data, meta)
}