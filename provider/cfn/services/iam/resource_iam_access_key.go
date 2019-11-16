// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMAccessKeyType string = "AWS::IAM::AccessKey"

func ResourceIAMAccessKey() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMAccessKeyExists,
		Read: resourceIAMAccessKeyRead,
		Create: resourceIAMAccessKeyCreate,
		Update: resourceIAMAccessKeyUpdate,
		Delete: resourceIAMAccessKeyDelete,
		CustomizeDiff: resourceIAMAccessKeyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"serial": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_name": {
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

func resourceIAMAccessKeyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIAMAccessKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMAccessKeyType, ResourceIAMAccessKey(), data, meta)
}

func resourceIAMAccessKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(iAMAccessKeyType, ResourceIAMAccessKey(), data, meta)
}

func resourceIAMAccessKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(iAMAccessKeyType, ResourceIAMAccessKey(), data, meta)
}

func resourceIAMAccessKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(iAMAccessKeyType, data, meta)
}

func resourceIAMAccessKeyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(iAMAccessKeyType, data, meta)
}
