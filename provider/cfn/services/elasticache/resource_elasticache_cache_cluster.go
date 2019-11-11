// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceElastiCacheCacheCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceElastiCacheCacheClusterCreate,
		Read:   resourceElastiCacheCacheClusterRead,
		Update: resourceElastiCacheCacheClusterUpdate,
		Delete: resourceElastiCacheCacheClusterDelete,

		Schema: map[string]*schema.Schema{
			"configuration_endpoint_address": {
				Type: schema.TypeString,
				Computed: true,
			},
			"configuration_endpoint_port": {
				Type: schema.TypeString,
				Computed: true,
			},
			"redis_endpoint_address": {
				Type: schema.TypeString,
				Computed: true,
			},
			"redis_endpoint_port": {
				Type: schema.TypeString,
				Computed: true,
			},
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
				ForceNew: true,
			},
			"cluster_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"engine": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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
				ForceNew: true,
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
				ForceNew: true,
				Set: schema.HashString,
			},
			"snapshot_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"snapshot_retention_limit": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"snapshot_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
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

func resourceElastiCacheCacheClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::CacheCluster", ResourceElastiCacheCacheCluster(), data, meta)
}

func resourceElastiCacheCacheClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::CacheCluster", ResourceElastiCacheCacheCluster(), data, meta)
}

func resourceElastiCacheCacheClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::CacheCluster", ResourceElastiCacheCacheCluster(), data, meta)
}

func resourceElastiCacheCacheClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::CacheCluster", data, meta)
}