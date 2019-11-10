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

func ResourceStackUserAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceStackUserAssociationCreate,
		Read:   resourceStackUserAssociationRead,
		Update: resourceStackUserAssociationUpdate,
		Delete: resourceStackUserAssociationDelete,

		Schema: map[string]*schema.Schema{
			"send_email_notification": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stack_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authentication_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceStackUserAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::StackUserAssociation", data, meta)
}

func resourceStackUserAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::StackUserAssociation", data, meta)
}

func resourceStackUserAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::StackUserAssociation", data, meta)
}

func resourceStackUserAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::StackUserAssociation", data, meta)
}