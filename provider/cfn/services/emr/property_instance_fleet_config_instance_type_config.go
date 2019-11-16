// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancetypeconfig.html

package emr

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var instanceFleetConfigInstanceTypeConfigProperties map[string]string = map[string]string{
	"bid_price": "BidPrice",
	"bid_price_as_percentage_of_on_demand_price": "BidPriceAsPercentageOfOnDemandPrice",
	"configurations": "Configurations",
	"ebs_configuration": "EbsConfiguration",
	"instance_type": "InstanceType",
	"weighted_capacity": "WeightedCapacity",
}

func propertyInstanceFleetConfigInstanceTypeConfig(extras...string) *schema.Resource {
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
			"bid_price": {
				Type: schema.TypeString,
				Optional: true,
			},
			"bid_price_as_percentage_of_on_demand_price": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyInstanceFleetConfigConfiguration(),
				Optional: true,
			},
			"ebs_configuration": {
				Type: schema.TypeList,
				Elem: propertyInstanceFleetConfigEbsConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"weighted_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
