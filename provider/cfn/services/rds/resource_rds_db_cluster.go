// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRDSDBCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceRDSDBClusterCreate,
		Read:   resourceRDSDBClusterRead,
		Update: resourceRDSDBClusterUpdate,
		Delete: resourceRDSDBClusterDelete,

		Schema: map[string]*schema.Schema{
			"endpoint_address": {
				Type: schema.TypeString,
				Computed: true,
			},
			"endpoint_port": {
				Type: schema.TypeString,
				Computed: true,
			},
			"read_endpoint_address": {
				Type: schema.TypeString,
				Computed: true,
			},
			"associated_roles": {
				Type: schema.TypeSet,
				Elem: propertyDBClusterDBClusterRole(),
				Optional: true,
			},
			"availability_zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
				Set: schema.HashString,
			},
			"backtrack_window": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"backup_retention_period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_cluster_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"database_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"deletion_protection": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"enable_cloudwatch_logs_exports": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"enable_iam_database_authentication": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"engine": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"engine_mode": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"master_user_password": {
				Type: schema.TypeString,
				Optional: true,
			},
			"master_username": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"preferred_backup_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"replication_source_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"restore_type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scaling_configuration": {
				Type: schema.TypeList,
				Elem: propertyDBClusterScalingConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"snapshot_identifier": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_db_cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"source_region": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"storage_encrypted": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"use_latest_restorable_time": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"vpc_security_group_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
		},
	}
}

func resourceRDSDBClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::DBCluster", ResourceRDSDBCluster(), data, meta)
}

func resourceRDSDBClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::DBCluster", ResourceRDSDBCluster(), data, meta)
}

func resourceRDSDBClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::DBCluster", ResourceRDSDBCluster(), data, meta)
}

func resourceRDSDBClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::DBCluster", data, meta)
}