// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-policy.html

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const s3BucketPolicyType string = "AWS::S3::BucketPolicy"

var s3BucketPolicyProperties map[string]string = map[string]string{
	"bucket": "Bucket",
	"policy_document": "PolicyDocument",
}

func ResourceS3BucketPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceS3BucketPolicyExists,
		Read: resourceS3BucketPolicyRead,
		Create: resourceS3BucketPolicyCreate,
		Update: resourceS3BucketPolicyUpdate,
		Delete: resourceS3BucketPolicyDelete,
		CustomizeDiff: resourceS3BucketPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type: schema.TypeString,
				Required: true,
			},
			"policy_document": {
				Type: schema.TypeMap,
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

func resourceS3BucketPolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceS3BucketPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(s3BucketPolicyType, ResourceS3BucketPolicy(), data, meta)
}

func resourceS3BucketPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(s3BucketPolicyType, ResourceS3BucketPolicy(), data, s3BucketPolicyProperties, meta)
}

func resourceS3BucketPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(s3BucketPolicyType, ResourceS3BucketPolicy(), data, s3BucketPolicyProperties, meta)
}

func resourceS3BucketPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(s3BucketPolicyType, data, meta)
}

func resourceS3BucketPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(s3BucketPolicyType, data, meta)
}
