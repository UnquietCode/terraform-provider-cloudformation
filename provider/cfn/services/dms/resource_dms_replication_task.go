// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDMSReplicationTask() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDMSReplicationTaskExists,
		Read: resourceDMSReplicationTaskRead,
		Create: resourceDMSReplicationTaskCreate,
		Update: resourceDMSReplicationTaskUpdate,
		Delete: resourceDMSReplicationTaskDelete,
		CustomizeDiff: resourceDMSReplicationTaskCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"replication_task_settings": {
				Type: schema.TypeString,
				Optional: true,
			},
			"table_mappings": {
				Type: schema.TypeString,
				Required: true,
			},
			"cdc_start_position": {
				Type: schema.TypeString,
				Optional: true,
			},
			"replication_task_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cdc_stop_position": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_endpoint_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"migration_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_endpoint_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"replication_instance_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"cdc_start_time": {
				Type: schema.TypeFloat,
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

func resourceDMSReplicationTaskExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDMSReplicationTaskRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::ReplicationTask", ResourceDMSReplicationTask(), data, meta)
}

func resourceDMSReplicationTaskCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::ReplicationTask", ResourceDMSReplicationTask(), data, meta)
}

func resourceDMSReplicationTaskUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::ReplicationTask", ResourceDMSReplicationTask(), data, meta)
}

func resourceDMSReplicationTaskDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::ReplicationTask", data, meta)
}

func resourceDMSReplicationTaskCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

