// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-partitioninput.html

package glue

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var partitionPartitionInputProperties map[string]string = map[string]string{
	"parameters": "Parameters",
	"storage_descriptor": "StorageDescriptor",
	"values": "Values",
}

func propertyPartitionPartitionInput(extras...string) *schema.Resource {
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
			"parameters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"storage_descriptor": {
				Type: schema.TypeList,
				Elem: propertyPartitionStorageDescriptor(),
				Optional: true,
				MaxItems: 1,
			},
			"values": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}
