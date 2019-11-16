// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamStackUserAssociationType string = "AWS::AppStream::StackUserAssociation"

var appStreamStackUserAssociationProperties map[string]string = map[string]string{
	"send_email_notification": "SendEmailNotification",
	"user_name": "UserName",
	"stack_name": "StackName",
	"authentication_type": "AuthenticationType",
}

func ResourceAppStreamStackUserAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppStreamStackUserAssociationExists,
		Read: resourceAppStreamStackUserAssociationRead,
		Create: resourceAppStreamStackUserAssociationCreate,
		Update: resourceAppStreamStackUserAssociationUpdate,
		Delete: resourceAppStreamStackUserAssociationDelete,
		CustomizeDiff: resourceAppStreamStackUserAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"send_email_notification": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stack_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"authentication_type": {
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

func resourceAppStreamStackUserAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppStreamStackUserAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appStreamStackUserAssociationType, ResourceAppStreamStackUserAssociation(), data, meta)
}

func resourceAppStreamStackUserAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appStreamStackUserAssociationType, ResourceAppStreamStackUserAssociation(), data, appStreamStackUserAssociationProperties, meta)
}

func resourceAppStreamStackUserAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appStreamStackUserAssociationType, ResourceAppStreamStackUserAssociation(), data, appStreamStackUserAssociationProperties, meta)
}

func resourceAppStreamStackUserAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appStreamStackUserAssociationType, data, meta)
}

func resourceAppStreamStackUserAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appStreamStackUserAssociationType, data, meta)
}
