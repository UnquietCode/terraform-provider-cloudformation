// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html

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
			"iam_auth_enabled": {
				Type: schema.TypeBool,
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
			"vpc_security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
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
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNeptuneDBClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Neptune::DBCluster", ResourceNeptuneDBCluster(), data, meta)
}

func resourceNeptuneDBClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Neptune::DBCluster", ResourceNeptuneDBCluster(), data, meta)
}

func resourceNeptuneDBClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Neptune::DBCluster", ResourceNeptuneDBCluster(), data, meta)
}

func resourceNeptuneDBClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Neptune::DBCluster", data, meta)
}