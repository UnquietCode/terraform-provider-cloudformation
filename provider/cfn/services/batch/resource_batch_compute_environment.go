// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-computeenvironment.html

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const batchComputeEnvironmentType string = "AWS::Batch::ComputeEnvironment"

var batchComputeEnvironmentProperties map[string]string = map[string]string{
	"type": "Type",
	"service_role": "ServiceRole",
	"compute_environment_name": "ComputeEnvironmentName",
	"compute_resources": "ComputeResources",
	"state": "State",
}

func ResourceBatchComputeEnvironment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceBatchComputeEnvironmentExists,
		Read: resourceBatchComputeEnvironmentRead,
		Create: resourceBatchComputeEnvironmentCreate,
		Update: resourceBatchComputeEnvironmentUpdate,
		Delete: resourceBatchComputeEnvironmentDelete,
		CustomizeDiff: resourceBatchComputeEnvironmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"service_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"compute_environment_name": {
				Type: schema.TypeString,
				Optional: true,
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
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBatchComputeEnvironmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceBatchComputeEnvironmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(batchComputeEnvironmentType, ResourceBatchComputeEnvironment(), data, meta)
}

func resourceBatchComputeEnvironmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(batchComputeEnvironmentType, ResourceBatchComputeEnvironment(), data, batchComputeEnvironmentProperties, meta)
}

func resourceBatchComputeEnvironmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(batchComputeEnvironmentType, ResourceBatchComputeEnvironment(), data, batchComputeEnvironmentProperties, meta)
}

func resourceBatchComputeEnvironmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(batchComputeEnvironmentType, data, meta)
}

func resourceBatchComputeEnvironmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(batchComputeEnvironmentType, data, meta)
}
