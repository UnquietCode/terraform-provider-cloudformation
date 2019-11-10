// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMAccessKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMAccessKeyCreate,
		Read:   resourceIAMAccessKeyRead,
		Update: resourceIAMAccessKeyUpdate,
		Delete: resourceIAMAccessKeyDelete,

		Schema: map[string]*schema.Schema{
			"serial": {
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIAMAccessKeyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::AccessKey", data, meta)
}

func resourceIAMAccessKeyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::AccessKey", data, meta)
}

func resourceIAMAccessKeyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::AccessKey", data, meta)
}

func resourceIAMAccessKeyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::AccessKey", data, meta)
}