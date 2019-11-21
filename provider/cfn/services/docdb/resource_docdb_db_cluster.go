// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html

package docdb

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const docDBDBClusterType string = "AWS::DocDB::DBCluster"

func ResourceDocDBDBCluster() *schema.Resource {
	return &schema.Resource{
		Read: resourceDocDBDBClusterRead,
		Create: resourceDocDBDBClusterCreate,
		Update: resourceDocDBDBClusterUpdate,
		Delete: resourceDocDBDBClusterDelete,
		CustomizeDiff: resourceDocDBDBClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"storage_encrypted": {
				Type: schema.TypeBool,
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
			"availability_zones": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"snapshot_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"db_cluster_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"backup_retention_period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"enable_cloudwatch_logs_exports": {
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
		},
	}
}

func resourceDocDBDBClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(docDBDBClusterType, ResourceDocDBDBCluster(), data, meta)
}

func resourceDocDBDBClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(docDBDBClusterType, ResourceDocDBDBCluster(), data, meta)
}

func resourceDocDBDBClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(docDBDBClusterType, ResourceDocDBDBCluster(), data, meta)
}

func resourceDocDBDBClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(docDBDBClusterType, data, meta)
}

func resourceDocDBDBClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(docDBDBClusterType, data, meta)
}
