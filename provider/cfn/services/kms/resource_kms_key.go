// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html

package kms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kMSKeyType string = "AWS::KMS::Key"

var kMSKeyProperties map[string]string = map[string]string{
	"description": "Description",
	"enable_key_rotation": "EnableKeyRotation",
	"enabled": "Enabled",
	"key_policy": "KeyPolicy",
	"key_usage": "KeyUsage",
	"pending_window_in_days": "PendingWindowInDays",
	"tags": "Tags",
}

func ResourceKMSKey() *schema.Resource {
	return &schema.Resource{
		Exists: resourceKMSKeyExists,
		Read: resourceKMSKeyRead,
		Create: resourceKMSKeyCreate,
		Update: resourceKMSKeyUpdate,
		Delete: resourceKMSKeyDelete,
		CustomizeDiff: resourceKMSKeyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable_key_rotation": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"key_policy": {
				Type: schema.TypeMap,
				Required: true,
			},
			"key_usage": {
				Type: schema.TypeString,
				Optional: true,
			},
			"pending_window_in_days": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceKMSKeyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKMSKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kMSKeyType, ResourceKMSKey(), data, meta)
}

func resourceKMSKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kMSKeyType, ResourceKMSKey(), data, kMSKeyProperties, meta)
}

func resourceKMSKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kMSKeyType, ResourceKMSKey(), data, kMSKeyProperties, meta)
}

func resourceKMSKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kMSKeyType, data, meta)
}

func resourceKMSKeyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kMSKeyType, data, meta)
}
