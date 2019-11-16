// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-parametergroup.html

package dax

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dAXParameterGroupType string = "AWS::DAX::ParameterGroup"

var dAXParameterGroupProperties map[string]string = map[string]string{
	"parameter_name_values": "ParameterNameValues",
	"description": "Description",
	"parameter_group_name": "ParameterGroupName",
}

func ResourceDAXParameterGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDAXParameterGroupExists,
		Read: resourceDAXParameterGroupRead,
		Create: resourceDAXParameterGroupCreate,
		Update: resourceDAXParameterGroupUpdate,
		Delete: resourceDAXParameterGroupDelete,
		CustomizeDiff: resourceDAXParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"parameter_name_values": {
				Type: schema.TypeMap,
				Optional: true,
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
			},
		},
	}
}

func resourceDAXParameterGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDAXParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dAXParameterGroupType, ResourceDAXParameterGroup(), data, meta)
}

func resourceDAXParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dAXParameterGroupType, ResourceDAXParameterGroup(), data, dAXParameterGroupProperties, meta)
}

func resourceDAXParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dAXParameterGroupType, ResourceDAXParameterGroup(), data, dAXParameterGroupProperties, meta)
}

func resourceDAXParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dAXParameterGroupType, data, meta)
}

func resourceDAXParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dAXParameterGroupType, data, meta)
}
