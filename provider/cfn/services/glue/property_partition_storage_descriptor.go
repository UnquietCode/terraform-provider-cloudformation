// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html

package glue

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyPartitionStorageDescriptor(extras...string) *schema.Resource {
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
			"stored_as_sub_directories": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"bucket_columns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"skewed_info": {
				Type: schema.TypeSet,
				Elem: propertyPartitionSkewedInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"input_format": {
				Type: schema.TypeString,
				Optional: true,
			},
			"number_of_buckets": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"output_format": {
				Type: schema.TypeString,
				Optional: true,
			},
			"columns": {
				Type: schema.TypeList,
				Elem: propertyPartitionColumn(),
				Optional: true,
			},
			"serde_info": {
				Type: schema.TypeSet,
				Elem: propertyPartitionSerdeInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"sort_columns": {
				Type: schema.TypeList,
				Elem: propertyPartitionOrder(),
				Optional: true,
			},
			"compressed": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"location": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
