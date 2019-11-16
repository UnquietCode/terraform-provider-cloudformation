// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMInstanceProfileType string = "AWS::IAM::InstanceProfile"

func ResourceIAMInstanceProfile() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMInstanceProfileExists,
		Read: resourceIAMInstanceProfileRead,
		Create: resourceIAMInstanceProfileCreate,
		Update: resourceIAMInstanceProfileUpdate,
		Delete: resourceIAMInstanceProfileDelete,
		CustomizeDiff: resourceIAMInstanceProfileCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"instance_profile_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"roles": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceIAMInstanceProfileExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIAMInstanceProfileRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMInstanceProfileType, ResourceIAMInstanceProfile(), data, meta)
}

func resourceIAMInstanceProfileCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(iAMInstanceProfileType, ResourceIAMInstanceProfile(), data, meta)
}

func resourceIAMInstanceProfileUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(iAMInstanceProfileType, ResourceIAMInstanceProfile(), data, meta)
}

func resourceIAMInstanceProfileDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(iAMInstanceProfileType, data, meta)
}

func resourceIAMInstanceProfileCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(iAMInstanceProfileType, data, meta)
}
