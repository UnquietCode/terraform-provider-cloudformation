// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceJobQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobQueueCreate,
		Read:   resourceJobQueueRead,
		Update: resourceJobQueueUpdate,
		Delete: resourceJobQueueDelete,

		Schema: map[string]*schema.Schema{
			"compute_environment_order": {
				Type: schema.TypeList,
				Elem: propertyJobQueueComputeEnvironmentOrder(),
				Required: true,
			},
			"priority": {
				Type: schema.TypeInt,
				Required: true,
			},
			"state": {
				Type: schema.TypeString,
				Required: false,
			},
			"job_queue_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceJobQueueCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Batch::JobQueue", data, meta)
}

func resourceJobQueueRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Batch::JobQueue", data, meta)
}

func resourceJobQueueUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Batch::JobQueue", data, meta)
}

func resourceJobQueueDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Batch::JobQueue", data, meta)
}