// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package transfer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTransferUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceTransferUserCreate,
		Read:   resourceTransferUserRead,
		Update: resourceTransferUserUpdate,
		Delete: resourceTransferUserDelete,

		Schema: map[string]*schema.Schema{
			"policy": {
				Type: schema.TypeString,
				Required: false,
			},
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"home_directory": {
				Type: schema.TypeString,
				Required: false,
			},
			"server_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_public_keys": {
				Type: schema.TypeList,
				Elem: propertyUserSshPublicKey(),
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceTransferUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Transfer::User", data, meta)
}

func resourceTransferUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Transfer::User", data, meta)
}

func resourceTransferUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Transfer::User", data, meta)
}

func resourceTransferUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Transfer::User", data, meta)
}