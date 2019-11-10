// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceDMSReplicationTaskCreate,
		Read:   resourceDMSReplicationTaskRead,
		Update: resourceDMSReplicationTaskUpdate,
		Delete: resourceDMSReplicationTaskDelete,

		Schema: map[string]*schema.Schema{
			"replication_task_settings": {
				Type: schema.TypeString,
				Required: false,
			},
			"table_mappings": {
				Type: schema.TypeString,
				Required: true,
			},
			"cdc_start_position": {
				Type: schema.TypeString,
				Required: false,
			},
			"replication_task_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"cdc_stop_position": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_endpoint_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"migration_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_endpoint_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"replication_instance_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
				ForceNew: true,
			},
			"cdc_start_time": {
				Type: schema.TypeFloat,
				Required: false,
			},
		},
	}
}

func resourceDMSReplicationTaskCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::ReplicationTask", data, meta)
}

func resourceDMSReplicationTaskRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::ReplicationTask", data, meta)
}

func resourceDMSReplicationTaskUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::ReplicationTask", data, meta)
}

func resourceDMSReplicationTaskDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::ReplicationTask", data, meta)
}