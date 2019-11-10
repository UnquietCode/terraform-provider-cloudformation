// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
			"replication_instance_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"engine_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"preferred_maintenance_window": {
				Type: schema.TypeString,
				Required: false,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Required: false,
			},
			"replication_subnet_group_identifier": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"allocated_storage": {
				Type: schema.TypeInt,
				Required: false,
			},
			"vpc_security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"allow_major_version_upgrade": {
				Type: schema.TypeBool,
				Required: false,
			},
			"replication_instance_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"publicly_accessible": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"multi_az": {
				Type: schema.TypeBool,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceDMSReplicationInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::ReplicationInstance", data, meta)
}

func resourceDMSReplicationInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::ReplicationInstance", data, meta)
}

func resourceDMSReplicationInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::ReplicationInstance", data, meta)
}

func resourceDMSReplicationInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::ReplicationInstance", data, meta)
}