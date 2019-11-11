// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceElastiCacheReplicationGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceElastiCacheReplicationGroupCreate,
		Read:   resourceElastiCacheReplicationGroupRead,
		Update: resourceElastiCacheReplicationGroupUpdate,
		Delete: resourceElastiCacheReplicationGroupDelete,

		Schema: map[string]*schema.Schema{
			"configuration_end_point_address": {
				Type: schema.TypeString,
				Computed: true,
			},
			"configuration_end_point_port": {
				Type: schema.TypeString,
				Computed: true,
			},
			"primary_end_point_address": {
				Type: schema.TypeString,
				Computed: true,
			},
			"primary_end_point_port": {
				Type: schema.TypeString,
				Computed: true,
			},
			"read_end_point_addresses": {
				Type: schema.TypeString,
				Computed: true,
			},
			"read_end_point_addresses_list": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"read_end_point_ports": {
				Type: schema.TypeString,
				Computed: true,
			},
			"read_end_point_ports_list": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"at_rest_encryption_enabled": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"auth_token": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"engine": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"preferred_cache_cluster_a_zs": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"replication_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"replication_group_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElastiCacheReplicationGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::ReplicationGroup", ResourceElastiCacheReplicationGroup(), data, meta)
}

func resourceElastiCacheReplicationGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::ReplicationGroup", ResourceElastiCacheReplicationGroup(), data, meta)
}

func resourceElastiCacheReplicationGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::ReplicationGroup", ResourceElastiCacheReplicationGroup(), data, meta)
}

func resourceElastiCacheReplicationGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::ReplicationGroup", data, meta)
}