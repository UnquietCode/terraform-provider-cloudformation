// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dax

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDAXCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceDAXClusterCreate,
		Read:   resourceDAXClusterRead,
		Update: resourceDAXClusterUpdate,
		Delete: resourceDAXClusterDelete,

		Schema: map[string]*schema.Schema{
			"sse_specification": {
				Type: schema.TypeList,
				Elem: propertyClusterSSESpecification(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"replication_factor": {
				Type: schema.TypeInt,
				Required: true,
			},
			"parameter_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"availability_zones": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"iam_role_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"notification_topic_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"node_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceDAXClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DAX::Cluster", data, meta)
}

func resourceDAXClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DAX::Cluster", data, meta)
}

func resourceDAXClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DAX::Cluster", data, meta)
}

func resourceDAXClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DAX::Cluster", data, meta)
}