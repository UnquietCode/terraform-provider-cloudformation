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

func ResourceLambdaAlias() *schema.Resource {
	return &schema.Resource{
		Create: resourceLambdaAliasCreate,
		Read:   resourceLambdaAliasRead,
		Update: resourceLambdaAliasUpdate,
		Delete: resourceLambdaAliasDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"function_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"function_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"routing_config": {
				Type: schema.TypeList,
				Elem: propertyAliasRoutingConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceLambdaAliasCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Lambda::Alias", data, meta)
}

func resourceLambdaAliasRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Lambda::Alias", data, meta)
}

func resourceLambdaAliasUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Lambda::Alias", data, meta)
}

func resourceLambdaAliasDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Lambda::Alias", data, meta)
}