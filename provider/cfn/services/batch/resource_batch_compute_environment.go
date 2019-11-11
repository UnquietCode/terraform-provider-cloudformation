// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBatchComputeEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceBatchComputeEnvironmentCreate,
		Read:   resourceBatchComputeEnvironmentRead,
		Update: resourceBatchComputeEnvironmentUpdate,
		Delete: resourceBatchComputeEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"compute_environment_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"compute_resources": {
				Type: schema.TypeList,
				Elem: propertyComputeEnvironmentComputeResources(),
				Optional: true,
				MaxItems: 1,
			},
			"state": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBatchComputeEnvironmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Batch::ComputeEnvironment", ResourceBatchComputeEnvironment(), data, meta)
}

func resourceBatchComputeEnvironmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Batch::ComputeEnvironment", ResourceBatchComputeEnvironment(), data, meta)
}

func resourceBatchComputeEnvironmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Batch::ComputeEnvironment", ResourceBatchComputeEnvironment(), data, meta)
}

func resourceBatchComputeEnvironmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Batch::ComputeEnvironment", data, meta)
}