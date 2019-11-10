// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppStreamStackUserAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppStreamStackUserAssociationCreate,
		Read:   resourceAppStreamStackUserAssociationRead,
		Update: resourceAppStreamStackUserAssociationUpdate,
		Delete: resourceAppStreamStackUserAssociationDelete,

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

func resourceAppStreamStackUserAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::StackUserAssociation", data, meta)
}

func resourceAppStreamStackUserAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::StackUserAssociation", data, meta)
}

func resourceAppStreamStackUserAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::StackUserAssociation", data, meta)
}

func resourceAppStreamStackUserAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::StackUserAssociation", data, meta)
}