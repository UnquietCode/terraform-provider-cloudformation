// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2FlowLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2FlowLogCreate,
		Read:   resourceEC2FlowLogRead,
		Delete: resourceEC2FlowLogDelete,

		Schema: map[string]*schema.Schema{
			"deliver_logs_permission_arn": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"log_destination": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"log_destination_type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"traffic_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2FlowLogCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::FlowLog", data, meta)
}

func resourceEC2FlowLogRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::FlowLog", data, meta)
}

func resourceEC2FlowLogUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::FlowLog", data, meta)
}

func resourceEC2FlowLogDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::FlowLog", data, meta)
}