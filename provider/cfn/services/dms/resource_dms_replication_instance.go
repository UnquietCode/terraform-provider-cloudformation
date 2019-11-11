// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDMSReplicationInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMSReplicationInstanceCreate,
		Read:   resourceDMSReplicationInstanceRead,
		Update: resourceDMSReplicationInstanceUpdate,
		Delete: resourceDMSReplicationInstanceDelete,

		Schema: map[string]*schema.Schema{
			"replication_instance_public_ip_addresses": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"replication_instance_private_ip_addresses": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"replication_instance_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kms_key_id": {
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
			"replication_subnet_group_identifier": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"allocated_storage": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"vpc_security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"allow_major_version_upgrade": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"replication_instance_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"publicly_accessible": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"multi_az": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDMSReplicationInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::ReplicationInstance", ResourceDMSReplicationInstance(), data, meta)
}

func resourceDMSReplicationInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::ReplicationInstance", ResourceDMSReplicationInstance(), data, meta)
}

func resourceDMSReplicationInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::ReplicationInstance", ResourceDMSReplicationInstance(), data, meta)
}

func resourceDMSReplicationInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::ReplicationInstance", data, meta)
}