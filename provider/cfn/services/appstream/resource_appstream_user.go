// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamUserType string = "AWS::AppStream::User"

func ResourceAppStreamUser() *schema.Resource {
	return &schema.Resource{
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

func resourceAppStreamUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(appStreamUserType, ResourceAppStreamUser(), data, meta)
}

func resourceAppStreamUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appStreamUserType, ResourceAppStreamUser(), data, meta)
}

func resourceAppStreamUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appStreamUserType, ResourceAppStreamUser(), data, meta)
}

func resourceAppStreamUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appStreamUserType, data, meta)
}

func resourceAppStreamUserCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appStreamUserType, data, meta)
}
