// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elastiCacheReplicationGroupType string = "AWS::ElastiCache::ReplicationGroup"

var elastiCacheReplicationGroupProperties map[string]string = map[string]string{
	"at_rest_encryption_enabled": "AtRestEncryptionEnabled",
	"auth_token": "AuthToken",
	"auto_minor_version_upgrade": "AutoMinorVersionUpgrade",
	"automatic_failover_enabled": "AutomaticFailoverEnabled",
	"cache_node_type": "CacheNodeType",
	"cache_parameter_group_name": "CacheParameterGroupName",
	"cache_security_group_names": "CacheSecurityGroupNames",
	"cache_subnet_group_name": "CacheSubnetGroupName",
	"engine": "Engine",
	"engine_version": "EngineVersion",
	"kms_key_id": "KmsKeyId",
	"node_group_configuration": "NodeGroupConfiguration",
	"notification_topic_arn": "NotificationTopicArn",
	"num_cache_clusters": "NumCacheClusters",
	"num_node_groups": "NumNodeGroups",
	"port": "Port",
	"preferred_cache_cluster_a_zs": "PreferredCacheClusterAZs",
	"preferred_maintenance_window": "PreferredMaintenanceWindow",
	"primary_cluster_id": "PrimaryClusterId",
	"replicas_per_node_group": "ReplicasPerNodeGroup",
	"replication_group_description": "ReplicationGroupDescription",
	"replication_group_id": "ReplicationGroupId",
	"security_group_ids": "SecurityGroupIds",
	"snapshot_arns": "SnapshotArns",
	"snapshot_name": "SnapshotName",
	"snapshot_retention_limit": "SnapshotRetentionLimit",
	"snapshot_window": "SnapshotWindow",
	"snapshotting_cluster_id": "SnapshottingClusterId",
	"tags": "Tags",
	"transit_encryption_enabled": "TransitEncryptionEnabled",
}

func ResourceElastiCacheReplicationGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElastiCacheReplicationGroupExists,
		Read: resourceElastiCacheReplicationGroupRead,
		Create: resourceElastiCacheReplicationGroupCreate,
		Update: resourceElastiCacheReplicationGroupUpdate,
		Delete: resourceElastiCacheReplicationGroupDelete,
		CustomizeDiff: resourceElastiCacheReplicationGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"at_rest_encryption_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"auth_token": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"automatic_failover_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"cache_node_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cache_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cache_security_group_names": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"cache_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"node_group_configuration": {
				Type: schema.TypeSet,
				Elem: propertyReplicationGroupNodeGroupConfiguration(),
				Optional: true,
			},
			"notification_topic_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"num_cache_clusters": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"num_node_groups": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"preferred_cache_cluster_a_zs": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"primary_cluster_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"replicas_per_node_group": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"replication_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"replication_group_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"snapshot_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"snapshot_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"snapshot_retention_limit": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"snapshot_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"snapshotting_cluster_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"transit_encryption_enabled": {
				Type: schema.TypeBool,
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

func resourceElastiCacheReplicationGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElastiCacheReplicationGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elastiCacheReplicationGroupType, ResourceElastiCacheReplicationGroup(), data, meta)
}

func resourceElastiCacheReplicationGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elastiCacheReplicationGroupType, ResourceElastiCacheReplicationGroup(), data, elastiCacheReplicationGroupProperties, meta)
}

func resourceElastiCacheReplicationGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elastiCacheReplicationGroupType, ResourceElastiCacheReplicationGroup(), data, elastiCacheReplicationGroupProperties, meta)
}

func resourceElastiCacheReplicationGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elastiCacheReplicationGroupType, data, meta)
}

func resourceElastiCacheReplicationGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elastiCacheReplicationGroupType, data, meta)
}
