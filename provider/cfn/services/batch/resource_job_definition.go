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

func ResourceJobDefinition() *schema.Resource {
	return &schema.Resource{
		Create: resourceJobDefinitionCreate,
		Read:   resourceJobDefinitionRead,
		Update: resourceJobDefinitionUpdate,
		Delete: resourceJobDefinitionDelete,

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

func resourceJobDefinitionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Batch::JobDefinition", data, meta)
}

func resourceJobDefinitionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Batch::JobDefinition", data, meta)
}

func resourceJobDefinitionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Batch::JobDefinition", data, meta)
}

func resourceJobDefinitionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Batch::JobDefinition", data, meta)
}