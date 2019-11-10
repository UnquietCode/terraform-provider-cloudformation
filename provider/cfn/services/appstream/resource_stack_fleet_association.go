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

func ResourceStackFleetAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceStackFleetAssociationCreate,
		Read:   resourceStackFleetAssociationRead,
		Update: resourceStackFleetAssociationUpdate,
		Delete: resourceStackFleetAssociationDelete,

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

func resourceStackFleetAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::StackFleetAssociation", data, meta)
}

func resourceStackFleetAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::StackFleetAssociation", data, meta)
}

func resourceStackFleetAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::StackFleetAssociation", data, meta)
}

func resourceStackFleetAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::StackFleetAssociation", data, meta)
}