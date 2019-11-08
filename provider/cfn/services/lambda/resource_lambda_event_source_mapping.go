// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLambdaEventSourceMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceLambdaEventSourceMappingCreate,
		Read:   resourceLambdaEventSourceMappingRead,
		Update: resourceLambdaEventSourceMappingUpdate,
		Delete: resourceLambdaEventSourceMappingDelete,

		Schema: map[string]*schema.Schema{
			"batch_size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"event_source_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"maximum_batching_window_in_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"starting_position": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceLambdaEventSourceMappingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::EventSourceMapping", data, meta)
}

func resourceLambdaEventSourceMappingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::EventSourceMapping", data, meta)
}

func resourceLambdaEventSourceMappingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::EventSourceMapping", data, meta)
}

func resourceLambdaEventSourceMappingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::EventSourceMapping", data, meta)
}