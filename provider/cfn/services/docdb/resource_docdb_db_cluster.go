// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html

package docdb

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDocDBDBCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceDocDBDBClusterCreate,
		Read:   resourceDocDBDBClusterRead,
		Update: resourceDocDBDBClusterUpdate,
		Delete: resourceDocDBDBClusterDelete,

		Schema: map[string]*schema.Schema{
			"cluster_resource_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"endpoint": {
				Type: schema.TypeString,
				Computed: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"read_endpoint": {
				Type: schema.TypeString,
				Computed: true,
			},
			"storage_encrypted": {
				Type: schema.TypeBool,
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
			"availability_zones": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"snapshot_identifier": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"preferred_backup_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"master_user_password": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"master_username": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"db_cluster_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"backup_retention_period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"enable_cloudwatch_logs_exports": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceDocDBDBClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DocDB::DBCluster", ResourceDocDBDBCluster(), data, meta)
}

func resourceDocDBDBClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DocDB::DBCluster", ResourceDocDBDBCluster(), data, meta)
}

func resourceDocDBDBClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DocDB::DBCluster", ResourceDocDBDBCluster(), data, meta)
}

func resourceDocDBDBClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DocDB::DBCluster", data, meta)
}