// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKMSKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceKMSKeyCreate,
		Read:   resourceKMSKeyRead,
		Update: resourceKMSKeyUpdate,
		Delete: resourceKMSKeyDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"enable_key_rotation": {
				Type: schema.TypeBool,
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"key_policy": {
				Type: schema.TypeMap,
				Required: true,
			},
			"key_usage": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"pending_window_in_days": {
				Type: schema.TypeInt,
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

func resourceKMSKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KMS::Key", data, meta)
}

func resourceKMSKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KMS::Key", data, meta)
}

func resourceKMSKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KMS::Key", data, meta)
}

func resourceKMSKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KMS::Key", data, meta)
}