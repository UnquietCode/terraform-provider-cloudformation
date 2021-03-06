// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html

package rds

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSDBInstanceType string = "AWS::RDS::DBInstance"

func DatasourceRDSDBInstance() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRDSDBInstanceRead,
		
		Schema: map[string]*schema.Schema{
			"allocated_storage": {
				Type: schema.TypeString,
				Optional: true,
			},
			"allow_major_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"associated_roles": {
				Type: schema.TypeSet,
				Elem: propertyDBInstanceDBInstanceRole(),
				Optional: true,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"backup_retention_period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"character_set_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"copy_tags_to_snapshot": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_instance_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"db_instance_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"db_snapshot_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"delete_automated_backups": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"deletion_protection": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"domain": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domain_iam_role_name": {
				Type: schema.TypeString,
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
			"enable_performance_insights": {
				Type: schema.TypeBool,
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
			"iops": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"license_model": {
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
			"monitoring_interval": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"monitoring_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"multi_az": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"option_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"performance_insights_kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"performance_insights_retention_period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"port": {
				Type: schema.TypeString,
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
			"processor_features": {
				Type: schema.TypeSet,
				Elem: propertyDBInstanceProcessorFeature(),
				Optional: true,
			},
			"promotion_tier": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"publicly_accessible": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"source_db_instance_identifier": {
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
			"storage_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"timezone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"use_default_processor_features": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"vpc_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceRDSDBInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBInstanceType, DatasourceRDSDBInstance(), data, meta)
}
