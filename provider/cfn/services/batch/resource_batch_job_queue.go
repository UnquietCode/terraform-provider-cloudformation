// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBatchJobQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceBatchJobQueueCreate,
		Read:   resourceBatchJobQueueRead,
		Update: resourceBatchJobQueueUpdate,
		Delete: resourceBatchJobQueueDelete,

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
				Optional: true,
			},
			"job_queue_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBatchJobQueueCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Batch::JobQueue", ResourceBatchJobQueue(), data, meta)
}

func resourceBatchJobQueueRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Batch::JobQueue", ResourceBatchJobQueue(), data, meta)
}

func resourceBatchJobQueueUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Batch::JobQueue", ResourceBatchJobQueue(), data, meta)
}

func resourceBatchJobQueueDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Batch::JobQueue", data, meta)
}