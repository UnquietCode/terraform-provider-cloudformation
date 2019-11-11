// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceNeptuneDBInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceNeptuneDBInstanceCreate,
		Read:   resourceNeptuneDBInstanceRead,
		Update: resourceNeptuneDBInstanceUpdate,
		Delete: resourceNeptuneDBInstanceDelete,

		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type: schema.TypeString,
				Computed: true,
			},
			"port": {
				Type: schema.TypeString,
				Computed: true,
			},
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
				ForceNew: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"db_instance_identifier": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_snapshot_identifier": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNeptuneDBInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Neptune::DBInstance", ResourceNeptuneDBInstance(), data, meta)
}

func resourceNeptuneDBInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Neptune::DBInstance", ResourceNeptuneDBInstance(), data, meta)
}

func resourceNeptuneDBInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Neptune::DBInstance", ResourceNeptuneDBInstance(), data, meta)
}

func resourceNeptuneDBInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Neptune::DBInstance", data, meta)
}