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

func ResourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

		Schema: map[string]*schema.Schema{
			"user_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"first_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"message_action": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"last_name": {
				Type: schema.TypeString,
				Required: false,
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

func resourceUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::User", data, meta)
}

func resourceUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::User", data, meta)
}

func resourceUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::User", data, meta)
}

func resourceUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::User", data, meta)
}