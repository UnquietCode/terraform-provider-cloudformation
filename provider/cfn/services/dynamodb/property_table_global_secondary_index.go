// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-gsi.html

package dynamodb

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var tableGlobalSecondaryIndexProperties map[string]string = map[string]string{
	"index_name": "IndexName",
	"key_schema": "KeySchema",
	"projection": "Projection",
	"provisioned_throughput": "ProvisionedThroughput",
}

func propertyTableGlobalSecondaryIndex(extras...string) *schema.Resource {
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
			"index_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"key_schema": {
				Type: schema.TypeSet,
				Elem: propertyTableKeySchema(),
				Required: true,
			},
			"projection": {
				Type: schema.TypeList,
				Elem: propertyTableProjection(),
				Required: true,
				MaxItems: 1,
			},
			"provisioned_throughput": {
				Type: schema.TypeList,
				Elem: propertyTableProvisionedThroughput(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
