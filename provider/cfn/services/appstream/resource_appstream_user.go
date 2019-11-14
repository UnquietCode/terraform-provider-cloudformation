// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppStreamUser() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppStreamUserExists,
		Read: resourceAppStreamUserRead,
		Create: resourceAppStreamUserCreate,
		Update: resourceAppStreamUserUpdate,
		Delete: resourceAppStreamUserDelete,
		CustomizeDiff: resourceAppStreamUserCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"first_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"message_action": {
				Type: schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceAppStreamUserExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppStreamUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::User", ResourceAppStreamUser(), data, meta)
}

func resourceAppStreamUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::User", ResourceAppStreamUser(), data, meta)
}

func resourceAppStreamUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::User", ResourceAppStreamUser(), data, meta)
}

func resourceAppStreamUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::User", data, meta)
}

func resourceAppStreamUserCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

