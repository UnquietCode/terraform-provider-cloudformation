// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMParameter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSMParameterCreate,
		Read:   resourceSSMParameterRead,
		Update: resourceSSMParameterUpdate,
		Delete: resourceSSMParameterDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"policies": {
				Type: schema.TypeString,
				Required: false,
			},
			"allowed_pattern": {
				Type: schema.TypeString,
				Required: false,
			},
			"tier": {
				Type: schema.TypeString,
				Required: false,
			},
			"value": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceSSMParameterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::Parameter", data, meta)
}

func resourceSSMParameterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::Parameter", data, meta)
}

func resourceSSMParameterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::Parameter", data, meta)
}

func resourceSSMParameterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::Parameter", data, meta)
}