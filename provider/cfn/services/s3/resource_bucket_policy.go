// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBucketPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceBucketPolicyCreate,
		Read:   resourceBucketPolicyRead,
		Update: resourceBucketPolicyUpdate,
		Delete: resourceBucketPolicyDelete,

		Schema: map[string]*schema.Schema{
			"bucket": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
		},
	}
}

func resourceBucketPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::S3::BucketPolicy", data, meta)
}

func resourceBucketPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::S3::BucketPolicy", data, meta)
}

func resourceBucketPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::S3::BucketPolicy", data, meta)
}

func resourceBucketPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::S3::BucketPolicy", data, meta)
}