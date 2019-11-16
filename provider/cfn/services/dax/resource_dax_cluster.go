// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-cluster.html

package dax

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dAXClusterType string = "AWS::DAX::Cluster"

var dAXClusterProperties map[string]string = map[string]string{
	"sse_specification": "SSESpecification",
	"description": "Description",
	"replication_factor": "ReplicationFactor",
	"parameter_group_name": "ParameterGroupName",
	"availability_zones": "AvailabilityZones",
	"iam_role_arn": "IAMRoleARN",
	"subnet_group_name": "SubnetGroupName",
	"preferred_maintenance_window": "PreferredMaintenanceWindow",
	"notification_topic_arn": "NotificationTopicARN",
	"security_group_ids": "SecurityGroupIds",
	"node_type": "NodeType",
	"cluster_name": "ClusterName",
	"tags": "Tags",
}

func ResourceDAXCluster() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDAXClusterExists,
		Read: resourceDAXClusterRead,
		Create: resourceDAXClusterCreate,
		Update: resourceDAXClusterUpdate,
		Delete: resourceDAXClusterDelete,
		CustomizeDiff: resourceDAXClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"sse_specification": {
				Type: schema.TypeList,
				Elem: propertyClusterSSESpecification(),
				Optional: true,
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
			},
			"subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"cluster_name": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceDAXClusterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDAXClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dAXClusterType, ResourceDAXCluster(), data, meta)
}

func resourceDAXClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dAXClusterType, ResourceDAXCluster(), data, dAXClusterProperties, meta)
}

func resourceDAXClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dAXClusterType, ResourceDAXCluster(), data, dAXClusterProperties, meta)
}

func resourceDAXClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dAXClusterType, data, meta)
}

func resourceDAXClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dAXClusterType, data, meta)
}
