// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
			"at_rest_encryption_enabled": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"auth_token": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Required: false,
			},
			"automatic_failover_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"cache_node_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"cache_parameter_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"cache_security_group_names": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"cache_subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"engine": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"node_group_configuration": {
				Type: schema.TypeSet,
				Elem: propertyReplicationGroupNodeGroupConfiguration(),
				Required: false,
			},
			"notification_topic_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"num_cache_clusters": {
				Type: schema.TypeInt,
				Required: false,
			},
			"num_node_groups": {
				Type: schema.TypeInt,
				Required: false,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"preferred_cache_cluster_a_zs": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"primary_cluster_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"replicas_per_node_group": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"replication_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"replication_group_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"security_group_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"snapshot_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"snapshot_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"snapshot_retention_limit": {
				Type: schema.TypeInt,
				Required: false,
			},
			"snapshot_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"snapshotting_cluster_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"transit_encryption_enabled": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceElastiCacheReplicationGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElastiCache::ReplicationGroup", data, meta)
}

func resourceElastiCacheReplicationGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElastiCache::ReplicationGroup", data, meta)
}

func resourceElastiCacheReplicationGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElastiCache::ReplicationGroup", data, meta)
}

func resourceElastiCacheReplicationGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElastiCache::ReplicationGroup", data, meta)
}