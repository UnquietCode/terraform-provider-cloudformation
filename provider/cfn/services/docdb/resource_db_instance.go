// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package docdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDBInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceDBInstanceCreate,
		Read:   resourceDBInstanceRead,
		Update: resourceDBInstanceUpdate,
		Delete: resourceDBInstanceDelete,

		Schema: map[string]*schema.Schema{
			"db_instance_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Required: false,
			},
			"db_instance_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceDBInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DocDB::DBInstance", data, meta)
}

func resourceDBInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DocDB::DBInstance", data, meta)
}

func resourceDBInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DocDB::DBInstance", data, meta)
}

func resourceDBInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DocDB::DBInstance", data, meta)
}