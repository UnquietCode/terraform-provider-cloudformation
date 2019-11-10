// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterJobFlowInstancesConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_master_security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"additional_slave_security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"core_instance_fleet": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceFleetConfig(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"core_instance_group": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceGroupConfig(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"ec2_key_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ec2_subnet_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ec2_subnet_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"emr_managed_master_security_group": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"emr_managed_slave_security_group": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"hadoop_version": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"keep_job_flow_alive_when_no_steps": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"master_instance_fleet": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceFleetConfig(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"master_instance_group": {
				Type: schema.TypeList,
				Elem: propertyClusterInstanceGroupConfig(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"placement": {
				Type: schema.TypeList,
				Elem: propertyClusterPlacementType(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"service_access_security_group": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"termination_protected": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}