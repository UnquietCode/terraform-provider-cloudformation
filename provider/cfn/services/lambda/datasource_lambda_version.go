// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-version.html

package lambda

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaVersionType string = "AWS::Lambda::Version"

func DatasourceLambdaVersion() *schema.Resource {
	return &schema.Resource{
		Read: datasourceLambdaVersionRead,
		
		Schema: map[string]*schema.Schema{
			"code_sha256": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"provisioned_concurrency_config": {
				Type: schema.TypeList,
				Elem: propertyVersionProvisionedConcurrencyConfiguration(),
				Optional: true,
				MaxItems: 1,
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

func datasourceLambdaVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaVersionType, DatasourceLambdaVersion(), data, meta)
}
