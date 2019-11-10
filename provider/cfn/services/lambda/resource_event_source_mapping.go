// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEventSourceMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventSourceMappingCreate,
		Read:   resourceEventSourceMappingRead,
		Update: resourceEventSourceMappingUpdate,
		Delete: resourceEventSourceMappingDelete,

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

func resourceEventSourceMappingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::EventSourceMapping", data, meta)
}

func resourceEventSourceMappingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::EventSourceMapping", data, meta)
}

func resourceEventSourceMappingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::EventSourceMapping", data, meta)
}

func resourceEventSourceMappingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::EventSourceMapping", data, meta)
}