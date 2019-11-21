// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-s3action.html

package ses

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyReceiptRuleS3Action(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
	    if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
	        count = i
	    }
	}
	
	if count >= 5 {
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"kms_key_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"topic_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"object_key_prefix": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
