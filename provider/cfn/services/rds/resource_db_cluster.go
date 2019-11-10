// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDBCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBClusterCreate,
		Read:   resourceDBClusterRead,
		Update: resourceDBClusterUpdate,
		Delete: resourceDBClusterDelete,

		Schema: map[string]*schema.Schema{
			"associated_roles": {
				Type: schema.TypeSet,
				Elem: propertyDBClusterDBClusterRole(),
				Required: false,
			},
			"availability_zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"backtrack_window": {
				Type: schema.TypeInt,
				Required: false,
			},
			"backup_retention_period": {
				Type: schema.TypeInt,
				Required: false,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"db_cluster_parameter_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"deletion_protection": {
				Type: schema.TypeBool,
				Required: false,
			},
			"enable_cloudwatch_logs_exports": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"enable_iam_database_authentication": {
				Type: schema.TypeBool,
				Required: false,
			},
			"engine": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"engine_mode": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"master_user_password": {
				Type: schema.TypeString,
				Required: false,
			},
			"master_username": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"preferred_backup_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"replication_source_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"restore_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"scaling_configuration": {
				Type: schema.TypeList,
				Elem: propertyDBClusterScalingConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"snapshot_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"source_db_cluster_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"source_region": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"storage_encrypted": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"use_latest_restorable_time": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"vpc_security_group_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
		},
	}
}

func resourceDBClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::DBCluster", data, meta)
}

func resourceDBClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::DBCluster", data, meta)
}

func resourceDBClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::DBCluster", data, meta)
}

func resourceDBClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::DBCluster", data, meta)
}