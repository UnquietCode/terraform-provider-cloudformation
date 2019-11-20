// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-parameter.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sSMParameterType string = "AWS::SSM::Parameter"

func ResourceSSMParameter() *schema.Resource {
	return &schema.Resource{
		Read: resourceSSMParameterRead,
		Create: resourceSSMParameterCreate,
		Update: resourceSSMParameterUpdate,
		Delete: resourceSSMParameterDelete,
		CustomizeDiff: resourceSSMParameterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"policies": {
				Type: schema.TypeString,
				Optional: true,
			},
			"allowed_pattern": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"value": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
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

func resourceSSMParameterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sSMParameterType, ResourceSSMParameter(), data, meta)
}

func resourceSSMParameterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sSMParameterType, ResourceSSMParameter(), data, meta)
}

func resourceSSMParameterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sSMParameterType, ResourceSSMParameter(), data, meta)
}

func resourceSSMParameterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sSMParameterType, data, meta)
}

func resourceSSMParameterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sSMParameterType, data, meta)
}
