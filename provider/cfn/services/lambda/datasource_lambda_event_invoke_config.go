// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-eventinvokeconfig.html

package lambda

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaEventInvokeConfigType string = "AWS::Lambda::EventInvokeConfig"

func DatasourceLambdaEventInvokeConfig() *schema.Resource {
	return &schema.Resource{
		Read: datasourceLambdaEventInvokeConfigRead,
		
		Schema: map[string]*schema.Schema{
			"function_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"maximum_retry_attempts": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"destination_config": {
				Type: schema.TypeList,
				Elem: propertyEventInvokeConfigDestinationConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"qualifier": {
				Type: schema.TypeString,
				Required: true,
			},
			"maximum_event_age_in_seconds": {
				Type: schema.TypeInt,
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

func datasourceLambdaEventInvokeConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaEventInvokeConfigType, DatasourceLambdaEventInvokeConfig(), data, meta)
}
