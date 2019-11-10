// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package redshift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRedshiftCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedshiftClusterCreate,
		Read:   resourceRedshiftClusterRead,
		Update: resourceRedshiftClusterUpdate,
		Delete: resourceRedshiftClusterDelete,

		Schema: map[string]*schema.Schema{
			"allow_version_upgrade": {
				Type: schema.TypeBool,
				Required: false,
			},
			"automated_snapshot_retention_period": {
				Type: schema.TypeInt,
				Required: false,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"cluster_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"cluster_parameter_group_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"cluster_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"cluster_subnet_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"cluster_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"cluster_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"db_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"elastic_ip": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"encrypted": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"hsm_client_certificate_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"hsm_configuration_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"iam_roles": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"logging_properties": {
				Type: schema.TypeList,
				Elem: propertyClusterLoggingProperties(),
				Required: false,
				MaxItems: 1,
			},
			"master_user_password": {
				Type: schema.TypeString,
				Required: true,
			},
			"master_username": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"node_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"number_of_nodes": {
				Type: schema.TypeInt,
				Required: false,
			},
			"owner_account": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"publicly_accessible": {
				Type: schema.TypeBool,
				Required: false,
			},
			"snapshot_cluster_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"snapshot_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"vpc_security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceRedshiftClusterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Redshift::Cluster", data, meta)
}

func resourceRedshiftClusterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Redshift::Cluster", data, meta)
}

func resourceRedshiftClusterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Redshift::Cluster", data, meta)
}

func resourceRedshiftClusterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Redshift::Cluster", data, meta)
}