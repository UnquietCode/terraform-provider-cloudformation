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
				Required: false,
			},
			"node_properties": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionNodeProperties(),
				Required: false,
				MaxItems: 1,
			},
			"timeout": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionTimeout(),
				Required: false,
				MaxItems: 1,
			},
			"container_properties": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionContainerProperties(),
				Required: false,
				MaxItems: 1,
			},
			"job_definition_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"retry_strategy": {
				Type: schema.TypeList,
				Elem: propertyJobDefinitionRetryStrategy(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceBatchJobDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Batch::JobDefinition", data, meta)
}

func resourceBatchJobDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Batch::JobDefinition", data, meta)
}

func resourceBatchJobDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Batch::JobDefinition", data, meta)
}

func resourceBatchJobDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Batch::JobDefinition", data, meta)
}