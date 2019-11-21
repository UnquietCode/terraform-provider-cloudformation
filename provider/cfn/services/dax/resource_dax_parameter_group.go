// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-parametergroup.html

package dax

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dAXParameterGroupType string = "AWS::DAX::ParameterGroup"

func ResourceDAXParameterGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceDAXParameterGroupRead,
		Create: resourceDAXParameterGroupCreate,
		Update: resourceDAXParameterGroupUpdate,
		Delete: resourceDAXParameterGroupDelete,
		CustomizeDiff: resourceDAXParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"parameter_name_values": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"parameter_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceDAXParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dAXParameterGroupType, ResourceDAXParameterGroup(), data, meta)
}

func resourceDAXParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dAXParameterGroupType, ResourceDAXParameterGroup(), data, meta)
}

func resourceDAXParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dAXParameterGroupType, ResourceDAXParameterGroup(), data, meta)
}

func resourceDAXParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dAXParameterGroupType, data, meta)
}

func resourceDAXParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dAXParameterGroupType, data, meta)
}
