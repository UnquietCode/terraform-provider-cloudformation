// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbinstance.html

package neptune

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const neptuneDBInstanceType string = "AWS::Neptune::DBInstance"

var neptuneDBInstanceProperties map[string]string = map[string]string{
	"db_parameter_group_name": "DBParameterGroupName",
	"db_instance_class": "DBInstanceClass",
	"allow_major_version_upgrade": "AllowMajorVersionUpgrade",
	"db_cluster_identifier": "DBClusterIdentifier",
	"availability_zone": "AvailabilityZone",
	"preferred_maintenance_window": "PreferredMaintenanceWindow",
	"auto_minor_version_upgrade": "AutoMinorVersionUpgrade",
	"db_subnet_group_name": "DBSubnetGroupName",
	"db_instance_identifier": "DBInstanceIdentifier",
	"db_snapshot_identifier": "DBSnapshotIdentifier",
	"tags": "Tags",
}

func ResourceNeptuneDBInstance() *schema.Resource {
	return &schema.Resource{
		Exists: resourceNeptuneDBInstanceExists,
		Read: resourceNeptuneDBInstanceRead,
		Create: resourceNeptuneDBInstanceCreate,
		Update: resourceNeptuneDBInstanceUpdate,
		Delete: resourceNeptuneDBInstanceDelete,
		CustomizeDiff: resourceNeptuneDBInstanceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"db_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_instance_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"allow_major_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_instance_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_snapshot_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNeptuneDBInstanceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceNeptuneDBInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(neptuneDBInstanceType, ResourceNeptuneDBInstance(), data, meta)
}

func resourceNeptuneDBInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(neptuneDBInstanceType, ResourceNeptuneDBInstance(), data, neptuneDBInstanceProperties, meta)
}

func resourceNeptuneDBInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(neptuneDBInstanceType, ResourceNeptuneDBInstance(), data, neptuneDBInstanceProperties, meta)
}

func resourceNeptuneDBInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(neptuneDBInstanceType, data, meta)
}

func resourceNeptuneDBInstanceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(neptuneDBInstanceType, data, meta)
}
