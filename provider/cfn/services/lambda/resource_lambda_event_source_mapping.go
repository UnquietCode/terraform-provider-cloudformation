// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaEventSourceMappingType string = "AWS::Lambda::EventSourceMapping"

var lambdaEventSourceMappingProperties map[string]string = map[string]string{
	"batch_size": "BatchSize",
	"enabled": "Enabled",
	"event_source_arn": "EventSourceArn",
	"function_name": "FunctionName",
	"maximum_batching_window_in_seconds": "MaximumBatchingWindowInSeconds",
	"starting_position": "StartingPosition",
}

func ResourceLambdaEventSourceMapping() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLambdaEventSourceMappingExists,
		Read: resourceLambdaEventSourceMappingRead,
		Create: resourceLambdaEventSourceMappingCreate,
		Update: resourceLambdaEventSourceMappingUpdate,
		Delete: resourceLambdaEventSourceMappingDelete,
		CustomizeDiff: resourceLambdaEventSourceMappingCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"batch_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"event_source_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"maximum_batching_window_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"starting_position": {
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

func resourceLambdaEventSourceMappingExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLambdaEventSourceMappingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaEventSourceMappingType, ResourceLambdaEventSourceMapping(), data, meta)
}

func resourceLambdaEventSourceMappingCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(lambdaEventSourceMappingType, ResourceLambdaEventSourceMapping(), data, lambdaEventSourceMappingProperties, meta)
}

func resourceLambdaEventSourceMappingUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(lambdaEventSourceMappingType, ResourceLambdaEventSourceMapping(), data, lambdaEventSourceMappingProperties, meta)
}

func resourceLambdaEventSourceMappingDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(lambdaEventSourceMappingType, data, meta)
}

func resourceLambdaEventSourceMappingCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(lambdaEventSourceMappingType, data, meta)
}
