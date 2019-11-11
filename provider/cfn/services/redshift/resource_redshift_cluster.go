// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html

package redshift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRedshiftCluster() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRedshiftClusterExists,
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
				Type: schema.TypeList,
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"vpc_security_group_ids": {
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

func resourceRedshiftClusterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRedshiftClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Redshift::Cluster", ResourceRedshiftCluster(), data, meta)
}

func resourceRedshiftClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Redshift::Cluster", ResourceRedshiftCluster(), data, meta)
}

func resourceRedshiftClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Redshift::Cluster", ResourceRedshiftCluster(), data, meta)
}

func resourceRedshiftClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Redshift::Cluster", data, meta)
}

func resourceRedshiftClusterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::Redshift::Cluster", data, meta)
}

