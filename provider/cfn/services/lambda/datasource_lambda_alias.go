// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html

package lambda

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const lambdaAliasType string = "AWS::Lambda::Alias"

func DatasourceLambdaAlias() *schema.Resource {
	return &schema.Resource{
		Read: datasourceLambdaAliasRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"function_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"provisioned_concurrency_config": {
				Type: schema.TypeList,
				Elem: propertyAliasProvisionedConcurrencyConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"routing_config": {
				Type: schema.TypeList,
				Elem: propertyAliasAliasRoutingConfiguration(),
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

func datasourceLambdaAliasRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(lambdaAliasType, DatasourceLambdaAlias(), data, meta)
}
