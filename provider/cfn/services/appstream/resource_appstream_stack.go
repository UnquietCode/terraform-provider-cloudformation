// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppStreamStack() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppStreamStackExists,
		Read:   resourceAppStreamStackRead,
		Create: resourceAppStreamStackCreate,
		Update: resourceAppStreamStackUpdate,
		Delete: resourceAppStreamStackDelete,
		CustomizeDiff: resourceAppStreamStackCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"application_settings": {
				Type: schema.TypeList,
				Elem: propertyStackApplicationSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"storage_connectors": {
				Type: schema.TypeList,
				Elem: propertyStackStorageConnector(),
				Optional: true,
			},
			"delete_storage_connectors": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"user_settings": {
				Type: schema.TypeList,
				Elem: propertyStackUserSetting(),
				Optional: true,
			},
			"attributes_to_delete": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"display_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"redirect_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"feedback_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppStreamStackExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppStreamStackRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::Stack", ResourceAppStreamStack(), data, meta)
}

func resourceAppStreamStackCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::Stack", ResourceAppStreamStack(), data, meta)
}

func resourceAppStreamStackUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::Stack", ResourceAppStreamStack(), data, meta)
}

func resourceAppStreamStackDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::Stack", data, meta)
}

func resourceAppStreamStackCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}