// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceFlowLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceFlowLogCreate,
		Read:   resourceFlowLogRead,
		Update: resourceFlowLogUpdate,
		Delete: resourceFlowLogDelete,

		Schema: map[string]*schema.Schema{
			"deliver_logs_permission_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"log_destination": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"log_destination_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Required: false,
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

func resourceFlowLogCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::FlowLog", data, meta)
}

func resourceFlowLogRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::FlowLog", data, meta)
}

func resourceFlowLogUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::FlowLog", data, meta)
}

func resourceFlowLogDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::FlowLog", data, meta)
}