// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html

package docdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const docDBDBInstanceType string = "AWS::DocDB::DBInstance"

var docDBDBInstanceProperties map[string]string = map[string]string{
	"db_instance_class": "DBInstanceClass",
	"db_cluster_identifier": "DBClusterIdentifier",
	"availability_zone": "AvailabilityZone",
	"preferred_maintenance_window": "PreferredMaintenanceWindow",
	"auto_minor_version_upgrade": "AutoMinorVersionUpgrade",
	"db_instance_identifier": "DBInstanceIdentifier",
	"tags": "Tags",
}

func ResourceDocDBDBInstance() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDocDBDBInstanceExists,
		Read: resourceDocDBDBInstanceRead,
		Create: resourceDocDBDBInstanceCreate,
		Update: resourceDocDBDBInstanceUpdate,
		Delete: resourceDocDBDBInstanceDelete,
		CustomizeDiff: resourceDocDBDBInstanceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"db_instance_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Required: true,
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
			"db_instance_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceDocDBDBInstanceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDocDBDBInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(docDBDBInstanceType, ResourceDocDBDBInstance(), data, meta)
}

func resourceDocDBDBInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(docDBDBInstanceType, ResourceDocDBDBInstance(), data, docDBDBInstanceProperties, meta)
}

func resourceDocDBDBInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(docDBDBInstanceType, ResourceDocDBDBInstance(), data, docDBDBInstanceProperties, meta)
}

func resourceDocDBDBInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(docDBDBInstanceType, data, meta)
}

func resourceDocDBDBInstanceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(docDBDBInstanceType, data, meta)
}
