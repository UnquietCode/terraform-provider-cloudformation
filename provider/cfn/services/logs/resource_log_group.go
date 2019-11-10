// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLogGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogGroupCreate,
		Read:   resourceLogGroupRead,
		Update: resourceLogGroupUpdate,
		Delete: resourceLogGroupDelete,

		Schema: map[string]*schema.Schema{
			"log_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"retention_in_days": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}

func resourceLogGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::LogGroup", data, meta)
}

func resourceLogGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::LogGroup", data, meta)
}

func resourceLogGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::LogGroup", data, meta)
}

func resourceLogGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::LogGroup", data, meta)
}