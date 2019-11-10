// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMInstanceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMInstanceProfileCreate,
		Read:   resourceIAMInstanceProfileRead,
		Update: resourceIAMInstanceProfileUpdate,
		Delete: resourceIAMInstanceProfileDelete,

		Schema: map[string]*schema.Schema{
			"instance_profile_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"path": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"roles": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceIAMInstanceProfileCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::InstanceProfile", data, meta)
}

func resourceIAMInstanceProfileRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::InstanceProfile", data, meta)
}

func resourceIAMInstanceProfileUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::InstanceProfile", data, meta)
}

func resourceIAMInstanceProfileDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::InstanceProfile", data, meta)
}