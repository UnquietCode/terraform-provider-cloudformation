// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-encryptioninfo.html

package msk

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterEncryptionInfo(extras...string) *schema.Resource {
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
			"encryption_at_rest": {
				Type: schema.TypeSet,
				Elem: propertyClusterEncryptionAtRest(),
				Optional: true,
				MaxItems: 1,
			},
			"encryption_in_transit": {
				Type: schema.TypeSet,
				Elem: propertyClusterEncryptionInTransit(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
