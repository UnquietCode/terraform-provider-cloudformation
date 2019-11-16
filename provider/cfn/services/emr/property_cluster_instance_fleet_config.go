// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html

package emr

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var clusterInstanceFleetConfigProperties map[string]string = map[string]string{
	"instance_type_configs": "InstanceTypeConfigs",
	"launch_specifications": "LaunchSpecifications",
	"name": "Name",
	"target_on_demand_capacity": "TargetOnDemandCapacity",
	"target_spot_capacity": "TargetSpotCapacity",
}

func propertyClusterInstanceFleetConfig(extras...string) *schema.Resource {
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
			"instance_type_configs": {
				Type: schema.TypeSet,
				Elem: propertyClusterInstanceTypeConfig(),
				Optional: true,
			},
			"launch_specifications": {
				Type: schema.TypeSet,
				Elem: propertyClusterInstanceFleetProvisioningSpecifications(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target_on_demand_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"target_spot_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
