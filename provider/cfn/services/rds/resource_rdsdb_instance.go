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

func ResourceRDSDBInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceRDSDBInstanceCreate,
		Read:   resourceRDSDBInstanceRead,
		Update: resourceRDSDBInstanceUpdate,
		Delete: resourceRDSDBInstanceDelete,

		Schema: map[string]*schema.Schema{
			"allocated_storage": {
				Type: schema.TypeString,
				Required: false,
			},
			"allow_major_version_upgrade": {
				Type: schema.TypeBool,
				Required: false,
			},
			"associated_roles": {
				Type: schema.TypeSet,
				Elem: propertyDBInstanceDBInstanceRole(),
				Required: false,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Required: false,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"backup_retention_period": {
				Type: schema.TypeInt,
				Required: false,
			},
			"character_set_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"copy_tags_to_snapshot": {
				Type: schema.TypeBool,
				Required: false,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"db_instance_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"db_instance_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"db_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"db_parameter_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"db_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"db_snapshot_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"delete_automated_backups": {
				Type: schema.TypeBool,
				Required: false,
			},
			"deletion_protection": {
				Type: schema.TypeBool,
				Required: false,
			},
			"domain": {
				Type: schema.TypeString,
				Required: false,
			},
			"domain_iam_role_name": {
				Type: schema.TypeString,
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
			"enable_performance_insights": {
				Type: schema.TypeBool,
				Required: false,
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
			"iops": {
				Type: schema.TypeInt,
				Required: false,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"license_model": {
				Type: schema.TypeString,
				Required: false,
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
			"monitoring_interval": {
				Type: schema.TypeInt,
				Required: false,
			},
			"monitoring_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"multi_az": {
				Type: schema.TypeBool,
				Required: false,
			},
			"option_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"performance_insights_kms_key_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"performance_insights_retention_period": {
				Type: schema.TypeInt,
				Required: false,
			},
			"port": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"preferred_backup_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"processor_features": {
				Type: schema.TypeSet,
				Elem: propertyDBInstanceProcessorFeature(),
				Required: false,
			},
			"promotion_tier": {
				Type: schema.TypeInt,
				Required: false,
			},
			"publicly_accessible": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"source_db_instance_identifier": {
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
			"storage_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"timezone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"use_default_processor_features": {
				Type: schema.TypeBool,
				Required: false,
			},
			"vpc_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceRDSDBInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::DBInstance", data, meta)
}

func resourceRDSDBInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::DBInstance", data, meta)
}

func resourceRDSDBInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::DBInstance", data, meta)
}

func resourceRDSDBInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::DBInstance", data, meta)
}