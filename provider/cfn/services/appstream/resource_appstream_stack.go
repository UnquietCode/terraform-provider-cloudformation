// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppStreamStack() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppStreamStackCreate,
		Read:   resourceAppStreamStackRead,
		Update: resourceAppStreamStackUpdate,
		Delete: resourceAppStreamStackDelete,

		Schema: map[string]*schema.Schema{
			"application_settings": {
				Type: schema.TypeList,
				Elem: propertyStackApplicationSettings(),
				Required: false,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"storage_connectors": {
				Type: schema.TypeList,
				Elem: propertyStackStorageConnector(),
				Required: false,
			},
			"delete_storage_connectors": {
				Type: schema.TypeBool,
				Required: false,
			},
			"user_settings": {
				Type: schema.TypeList,
				Elem: propertyStackUserSetting(),
				Required: false,
			},
			"attributes_to_delete": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"display_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"redirect_url": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"feedback_url": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceAppStreamStackCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::Stack", data, meta)
}

func resourceAppStreamStackRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::Stack", data, meta)
}

func resourceAppStreamStackUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::Stack", data, meta)
}

func resourceAppStreamStackDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::Stack", data, meta)
}