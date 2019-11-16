// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2FlowLogType string = "AWS::EC2::FlowLog"

var eC2FlowLogProperties map[string]string = map[string]string{
	"deliver_logs_permission_arn": "DeliverLogsPermissionArn",
	"log_destination": "LogDestination",
	"log_destination_type": "LogDestinationType",
	"log_group_name": "LogGroupName",
	"resource_id": "ResourceId",
	"resource_type": "ResourceType",
	"traffic_type": "TrafficType",
}

func ResourceEC2FlowLog() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2FlowLogExists,
		Read: resourceEC2FlowLogRead,
		Create: resourceEC2FlowLogCreate,
		Update: resourceEC2FlowLogUpdate,
		Delete: resourceEC2FlowLogDelete,
		CustomizeDiff: resourceEC2FlowLogCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"deliver_logs_permission_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"log_destination": {
				Type: schema.TypeString,
				Optional: true,
			},
			"log_destination_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"traffic_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2FlowLogExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2FlowLogRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2FlowLogType, ResourceEC2FlowLog(), data, meta)
}

func resourceEC2FlowLogCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2FlowLogType, ResourceEC2FlowLog(), data, eC2FlowLogProperties, meta)
}

func resourceEC2FlowLogUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2FlowLogType, ResourceEC2FlowLog(), data, eC2FlowLogProperties, meta)
}

func resourceEC2FlowLogDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2FlowLogType, data, meta)
}

func resourceEC2FlowLogCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2FlowLogType, data, meta)
}
