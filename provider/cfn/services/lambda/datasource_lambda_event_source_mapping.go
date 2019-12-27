// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventsourcemapping.html

package lambda

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaEventSourceMappingType string = "AWS::Lambda::EventSourceMapping"

func DatasourceLambdaEventSourceMapping() *schema.Resource {
	return &schema.Resource{
		Read: datasourceLambdaEventSourceMappingRead,
		
		Schema: map[string]*schema.Schema{
			"batch_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"bisect_batch_on_function_error": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"destination_config": {
				Type: schema.TypeList,
				Elem: propertyEventSourceMappingDestinationConfig(),
				Optional: true,
				MaxItems: 1,
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
			"maximum_record_age_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"maximum_retry_attempts": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"parallelization_factor": {
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceLambdaEventSourceMappingRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaEventSourceMappingType, DatasourceLambdaEventSourceMapping(), data, meta)
}
