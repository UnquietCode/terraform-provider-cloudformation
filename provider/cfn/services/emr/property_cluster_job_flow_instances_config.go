// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html

package emr

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterJobFlowInstancesConfig(extras...string) *schema.Resource {
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
			"additional_master_security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"additional_slave_security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"core_instance_fleet": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceFleetConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"core_instance_group": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceGroupConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"ec2_key_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ec2_subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ec2_subnet_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"emr_managed_master_security_group": {
				Type: schema.TypeString,
				Optional: true,
			},
			"emr_managed_slave_security_group": {
				Type: schema.TypeString,
				Optional: true,
			},
			"hadoop_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"keep_job_flow_alive_when_no_steps": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"master_instance_fleet": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceFleetConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"master_instance_group": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceGroupConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"placement": {
				Type: schema.TypeList,
				Elem: propertyClusterPlacementType(),
				Optional: true,
				MaxItems: 1,
			},
			"service_access_security_group": {
				Type: schema.TypeString,
				Optional: true,
			},
			"termination_protected": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}

