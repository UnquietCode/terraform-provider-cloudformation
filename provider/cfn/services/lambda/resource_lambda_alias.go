// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-alias.html

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLambdaAlias() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLambdaAliasExists,
		Read: resourceLambdaAliasRead,
		Create: resourceLambdaAliasCreate,
		Update: resourceLambdaAliasUpdate,
		Delete: resourceLambdaAliasDelete,
		CustomizeDiff: resourceLambdaAliasCustomizeDiff,
		
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
			},
		},
	}
}

func resourceLambdaAliasExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLambdaAliasRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::Alias", ResourceLambdaAlias(), data, meta)
}

func resourceLambdaAliasCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::Alias", ResourceLambdaAlias(), data, meta)
}

func resourceLambdaAliasUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::Alias", ResourceLambdaAlias(), data, meta)
}

func resourceLambdaAliasDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::Alias", data, meta)
}

func resourceLambdaAliasCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
