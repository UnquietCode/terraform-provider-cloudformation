// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package neptune

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceNeptuneDBCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceNeptuneDBClusterCreate,
		Read:   resourceNeptuneDBClusterRead,
		Update: resourceNeptuneDBClusterUpdate,
		Delete: resourceNeptuneDBClusterDelete,

		Schema: map[string]*schema.Schema{
			"storage_encrypted": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"availability_zones": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"snapshot_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"iam_auth_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"preferred_backup_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"vpc_security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"db_cluster_parameter_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"backup_retention_period": {
				Type: schema.TypeInt,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"enable_cloudwatch_logs_exports": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceNeptuneDBClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Neptune::DBCluster", data, meta)
}

func resourceNeptuneDBClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Neptune::DBCluster", data, meta)
}

func resourceNeptuneDBClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Neptune::DBCluster", data, meta)
}

func resourceNeptuneDBClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Neptune::DBCluster", data, meta)
}