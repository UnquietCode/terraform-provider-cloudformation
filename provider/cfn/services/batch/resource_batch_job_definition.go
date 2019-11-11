// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceBatchJobDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceBatchJobDefinitionCreate,
		Read:   resourceBatchJobDefinitionRead,
		Update: resourceBatchJobDefinitionUpdate,
		Delete: resourceBatchJobDefinitionDelete,

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
				Type: schema.TypeList,
				Elem: propertyJobDefinitionNodeProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"timeout": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionTimeout(),
				Optional: true,
				MaxItems: 1,
			},
			"container_properties": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionContainerProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"job_definition_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"retry_strategy": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionRetryStrategy(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBatchJobDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Batch::JobDefinition", ResourceBatchJobDefinition(), data, meta)
}

func resourceBatchJobDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Batch::JobDefinition", ResourceBatchJobDefinition(), data, meta)
}

func resourceBatchJobDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Batch::JobDefinition", ResourceBatchJobDefinition(), data, meta)
}

func resourceBatchJobDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Batch::JobDefinition", data, meta)
}