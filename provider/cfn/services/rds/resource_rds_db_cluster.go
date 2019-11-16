// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const rDSDBClusterType string = "AWS::RDS::DBCluster"

var rDSDBClusterProperties map[string]string = map[string]string{
	"associated_roles": "AssociatedRoles",
	"availability_zones": "AvailabilityZones",
	"backtrack_window": "BacktrackWindow",
	"backup_retention_period": "BackupRetentionPeriod",
	"db_cluster_identifier": "DBClusterIdentifier",
	"db_cluster_parameter_group_name": "DBClusterParameterGroupName",
	"db_subnet_group_name": "DBSubnetGroupName",
	"database_name": "DatabaseName",
	"deletion_protection": "DeletionProtection",
	"enable_cloudwatch_logs_exports": "EnableCloudwatchLogsExports",
	"enable_iam_database_authentication": "EnableIAMDatabaseAuthentication",
	"engine": "Engine",
	"engine_mode": "EngineMode",
	"engine_version": "EngineVersion",
	"kms_key_id": "KmsKeyId",
	"master_user_password": "MasterUserPassword",
	"master_username": "MasterUsername",
	"port": "Port",
	"preferred_backup_window": "PreferredBackupWindow",
	"preferred_maintenance_window": "PreferredMaintenanceWindow",
	"replication_source_identifier": "ReplicationSourceIdentifier",
	"restore_type": "RestoreType",
	"scaling_configuration": "ScalingConfiguration",
	"snapshot_identifier": "SnapshotIdentifier",
	"source_db_cluster_identifier": "SourceDBClusterIdentifier",
	"source_region": "SourceRegion",
	"storage_encrypted": "StorageEncrypted",
	"tags": "Tags",
	"use_latest_restorable_time": "UseLatestRestorableTime",
	"vpc_security_group_ids": "VpcSecurityGroupIds",
}

func ResourceRDSDBCluster() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRDSDBClusterExists,
		Read: resourceRDSDBClusterRead,
		Create: resourceRDSDBClusterCreate,
		Update: resourceRDSDBClusterUpdate,
		Delete: resourceRDSDBClusterDelete,
		CustomizeDiff: resourceRDSDBClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"associated_roles": {
				Type: schema.TypeSet,
				Elem: propertyDBClusterDBClusterRole(),
				Optional: true,
			},
			"availability_zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
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
			},
			"db_cluster_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"database_name": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"engine_mode": {
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
			"master_user_password": {
				Type: schema.TypeString,
				Optional: true,
			},
			"master_username": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"source_db_cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_region": {
				Type: schema.TypeString,
				Optional: true,
			},
			"storage_encrypted": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"use_latest_restorable_time": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"vpc_security_group_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRDSDBClusterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRDSDBClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBClusterType, ResourceRDSDBCluster(), data, meta)
}

func resourceRDSDBClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rDSDBClusterType, ResourceRDSDBCluster(), data, rDSDBClusterProperties, meta)
}

func resourceRDSDBClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rDSDBClusterType, ResourceRDSDBCluster(), data, rDSDBClusterProperties, meta)
}

func resourceRDSDBClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rDSDBClusterType, data, meta)
}

func resourceRDSDBClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rDSDBClusterType, data, meta)
}
