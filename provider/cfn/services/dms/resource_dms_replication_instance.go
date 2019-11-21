// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html

package dms

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dMSReplicationInstanceType string = "AWS::DMS::ReplicationInstance"

func ResourceDMSReplicationInstance() *schema.Resource {
	return &schema.Resource{
		Read: resourceDMSReplicationInstanceRead,
		Create: resourceDMSReplicationInstanceCreate,
		Update: resourceDMSReplicationInstanceUpdate,
		Delete: resourceDMSReplicationInstanceDelete,
		CustomizeDiff: resourceDMSReplicationInstanceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
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
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"multi_az": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceDMSReplicationInstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dMSReplicationInstanceType, ResourceDMSReplicationInstance(), data, meta)
}

func resourceDMSReplicationInstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dMSReplicationInstanceType, ResourceDMSReplicationInstance(), data, meta)
}

func resourceDMSReplicationInstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dMSReplicationInstanceType, ResourceDMSReplicationInstance(), data, meta)
}

func resourceDMSReplicationInstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dMSReplicationInstanceType, data, meta)
}

func resourceDMSReplicationInstanceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dMSReplicationInstanceType, data, meta)
}
