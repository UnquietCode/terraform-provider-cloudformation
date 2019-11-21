// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html

package neptune

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const neptuneDBClusterType string = "AWS::Neptune::DBCluster"

func ResourceNeptuneDBCluster() *schema.Resource {
	return &schema.Resource{
		Read: resourceNeptuneDBClusterRead,
		Create: resourceNeptuneDBClusterCreate,
		Update: resourceNeptuneDBClusterUpdate,
		Delete: resourceNeptuneDBClusterDelete,
		CustomizeDiff: resourceNeptuneDBClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"storage_encrypted": {
				Type: schema.TypeBool,
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
			"iam_auth_enabled": {
				Type: schema.TypeBool,
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

func resourceNeptuneDBClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(neptuneDBClusterType, ResourceNeptuneDBCluster(), data, meta)
}

func resourceNeptuneDBClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(neptuneDBClusterType, ResourceNeptuneDBCluster(), data, meta)
}

func resourceNeptuneDBClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(neptuneDBClusterType, ResourceNeptuneDBCluster(), data, meta)
}

func resourceNeptuneDBClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(neptuneDBClusterType, data, meta)
}

func resourceNeptuneDBClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(neptuneDBClusterType, data, meta)
}
