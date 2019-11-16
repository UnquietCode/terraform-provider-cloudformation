// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const batchJobDefinitionType string = "AWS::Batch::JobDefinition"

var batchJobDefinitionProperties map[string]string = map[string]string{
	"type": "Type",
	"parameters": "Parameters",
	"node_properties": "NodeProperties",
	"timeout": "Timeout",
	"container_properties": "ContainerProperties",
	"job_definition_name": "JobDefinitionName",
	"retry_strategy": "RetryStrategy",
}

func ResourceBatchJobDefinition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceBatchJobDefinitionExists,
		Read: resourceBatchJobDefinitionRead,
		Create: resourceBatchJobDefinitionCreate,
		Update: resourceBatchJobDefinitionUpdate,
		Delete: resourceBatchJobDefinitionDelete,
		CustomizeDiff: resourceBatchJobDefinitionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"node_properties": {
				Type: schema.TypeSet,
				Elem: propertyJobDefinitionNodeProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"timeout": {
				Type: schema.TypeSet,
				Elem: propertyJobDefinitionTimeout(),
				Optional: true,
				MaxItems: 1,
			},
			"container_properties": {
				Type: schema.TypeSet,
				Elem: propertyJobDefinitionContainerProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"job_definition_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"retry_strategy": {
				Type: schema.TypeSet,
				Elem: propertyJobDefinitionRetryStrategy(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBatchJobDefinitionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceBatchJobDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(batchJobDefinitionType, ResourceBatchJobDefinition(), data, meta)
}

func resourceBatchJobDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(batchJobDefinitionType, ResourceBatchJobDefinition(), data, batchJobDefinitionProperties, meta)
}

func resourceBatchJobDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(batchJobDefinitionType, ResourceBatchJobDefinition(), data, batchJobDefinitionProperties, meta)
}

func resourceBatchJobDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(batchJobDefinitionType, data, meta)
}

func resourceBatchJobDefinitionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(batchJobDefinitionType, data, meta)
}
