// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationtask.html

package dms

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dMSReplicationTaskType string = "AWS::DMS::ReplicationTask"

func DatasourceDMSReplicationTask() *schema.Resource {
	return &schema.Resource{
		Read: datasourceDMSReplicationTaskRead,
		
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
			"tags": misc.PropertyTags(),
			"cdc_start_time": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceDMSReplicationTaskRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dMSReplicationTaskType, DatasourceDMSReplicationTask(), data, meta)
}
