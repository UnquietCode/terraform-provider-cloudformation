// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html

package elasticache

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var replicationGroupNodeGroupConfigurationProperties map[string]string = map[string]string{
	"node_group_id": "NodeGroupId",
	"primary_availability_zone": "PrimaryAvailabilityZone",
	"replica_availability_zones": "ReplicaAvailabilityZones",
	"replica_count": "ReplicaCount",
	"slots": "Slots",
}

func propertyReplicationGroupNodeGroupConfiguration(extras...string) *schema.Resource {
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
			"node_group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"primary_availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"replica_availability_zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"replica_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"slots": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
