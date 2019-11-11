// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html

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
			"cluster_discovery_endpoint": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"sse_specification": {
				Type: schema.TypeList,
				Elem: propertyClusterSSESpecification(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"replication_factor": {
				Type: schema.TypeInt,
				Required: true,
			},
			"parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"availability_zones": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"iam_role_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_topic_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"node_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDAXClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DAX::Cluster", ResourceDAXCluster(), data, meta)
}

func resourceDAXClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DAX::Cluster", ResourceDAXCluster(), data, meta)
}

func resourceDAXClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DAX::Cluster", ResourceDAXCluster(), data, meta)
}

func resourceDAXClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DAX::Cluster", data, meta)
}