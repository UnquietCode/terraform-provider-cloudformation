// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cache-cluster.html

package elasticache

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elastiCacheCacheClusterType string = "AWS::ElastiCache::CacheCluster"

var elastiCacheCacheClusterProperties map[string]string = map[string]string{
	"az_mode": "AZMode",
	"auto_minor_version_upgrade": "AutoMinorVersionUpgrade",
	"cache_node_type": "CacheNodeType",
	"cache_parameter_group_name": "CacheParameterGroupName",
	"cache_security_group_names": "CacheSecurityGroupNames",
	"cache_subnet_group_name": "CacheSubnetGroupName",
	"cluster_name": "ClusterName",
	"engine": "Engine",
	"engine_version": "EngineVersion",
	"notification_topic_arn": "NotificationTopicArn",
	"num_cache_nodes": "NumCacheNodes",
	"port": "Port",
	"preferred_availability_zone": "PreferredAvailabilityZone",
	"preferred_availability_zones": "PreferredAvailabilityZones",
	"preferred_maintenance_window": "PreferredMaintenanceWindow",
	"snapshot_arns": "SnapshotArns",
	"snapshot_name": "SnapshotName",
	"snapshot_retention_limit": "SnapshotRetentionLimit",
	"snapshot_window": "SnapshotWindow",
	"tags": "Tags",
	"vpc_security_group_ids": "VpcSecurityGroupIds",
}

func ResourceElastiCacheCacheCluster() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElastiCacheCacheClusterExists,
		Read: resourceElastiCacheCacheClusterRead,
		Create: resourceElastiCacheCacheClusterCreate,
		Update: resourceElastiCacheCacheClusterUpdate,
		Delete: resourceElastiCacheCacheClusterDelete,
		CustomizeDiff: resourceElastiCacheCacheClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"az_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"cache_node_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"cache_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cache_security_group_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"cache_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine": {
				Type: schema.TypeString,
				Required: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"notification_topic_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"num_cache_nodes": {
				Type: schema.TypeInt,
				Required: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"preferred_availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"preferred_availability_zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
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
			"tags": misc.PropertyTags(),
			"vpc_security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceElastiCacheCacheClusterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElastiCacheCacheClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elastiCacheCacheClusterType, ResourceElastiCacheCacheCluster(), data, meta)
}

func resourceElastiCacheCacheClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elastiCacheCacheClusterType, ResourceElastiCacheCacheCluster(), data, elastiCacheCacheClusterProperties, meta)
}

func resourceElastiCacheCacheClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elastiCacheCacheClusterType, ResourceElastiCacheCacheCluster(), data, elastiCacheCacheClusterProperties, meta)
}

func resourceElastiCacheCacheClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elastiCacheCacheClusterType, data, meta)
}

func resourceElastiCacheCacheClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elastiCacheCacheClusterType, data, meta)
}
