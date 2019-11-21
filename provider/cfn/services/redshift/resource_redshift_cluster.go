// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html

package redshift

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const redshiftClusterType string = "AWS::Redshift::Cluster"

func ResourceRedshiftCluster() *schema.Resource {
	return &schema.Resource{
		Read: resourceRedshiftClusterRead,
		Create: resourceRedshiftClusterCreate,
		Update: resourceRedshiftClusterUpdate,
		Delete: resourceRedshiftClusterDelete,
		CustomizeDiff: resourceRedshiftClusterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"allow_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"automated_snapshot_retention_period": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cluster_parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cluster_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"cluster_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cluster_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"cluster_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"elastic_ip": {
				Type: schema.TypeString,
				Optional: true,
			},
			"encrypted": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"hsm_client_certificate_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"hsm_configuration_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"iam_roles": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logging_properties": {
				Type: schema.TypeSet,
				Elem: propertyClusterLoggingProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"master_user_password": {
				Type: schema.TypeString,
				Required: true,
			},
			"master_username": {
				Type: schema.TypeString,
				Required: true,
			},
			"node_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"number_of_nodes": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"owner_account": {
				Type: schema.TypeString,
				Optional: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Optional: true,
			},
			"publicly_accessible": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"snapshot_cluster_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"snapshot_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"vpc_security_group_ids": {
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

func resourceRedshiftClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(redshiftClusterType, ResourceRedshiftCluster(), data, meta)
}

func resourceRedshiftClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(redshiftClusterType, ResourceRedshiftCluster(), data, meta)
}

func resourceRedshiftClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(redshiftClusterType, ResourceRedshiftCluster(), data, meta)
}

func resourceRedshiftClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(redshiftClusterType, data, meta)
}

func resourceRedshiftClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(redshiftClusterType, data, meta)
}
