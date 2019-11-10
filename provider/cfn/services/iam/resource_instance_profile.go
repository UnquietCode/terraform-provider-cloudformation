// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceInstanceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceProfileCreate,
		Read:   resourceInstanceProfileRead,
		Update: resourceInstanceProfileUpdate,
		Delete: resourceInstanceProfileDelete,

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

func resourceInstanceProfileCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::InstanceProfile", data, meta)
}

func resourceInstanceProfileRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::InstanceProfile", data, meta)
}

func resourceInstanceProfileUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::InstanceProfile", data, meta)
}

func resourceInstanceProfileDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::InstanceProfile", data, meta)
}