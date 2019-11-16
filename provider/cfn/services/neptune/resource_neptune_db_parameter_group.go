// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbparametergroup.html

package neptune

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const neptuneDBParameterGroupType string = "AWS::Neptune::DBParameterGroup"

var neptuneDBParameterGroupProperties map[string]string = map[string]string{
	"description": "Description",
	"parameters": "Parameters",
	"family": "Family",
	"tags": "Tags",
	"name": "Name",
}

func ResourceNeptuneDBParameterGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceNeptuneDBParameterGroupExists,
		Read: resourceNeptuneDBParameterGroupRead,
		Create: resourceNeptuneDBParameterGroupCreate,
		Update: resourceNeptuneDBParameterGroupUpdate,
		Delete: resourceNeptuneDBParameterGroupDelete,
		CustomizeDiff: resourceNeptuneDBParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Required: true,
			},
			"family": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
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

func resourceNeptuneDBParameterGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceNeptuneDBParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(neptuneDBParameterGroupType, ResourceNeptuneDBParameterGroup(), data, meta)
}

func resourceNeptuneDBParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(neptuneDBParameterGroupType, ResourceNeptuneDBParameterGroup(), data, neptuneDBParameterGroupProperties, meta)
}

func resourceNeptuneDBParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(neptuneDBParameterGroupType, ResourceNeptuneDBParameterGroup(), data, neptuneDBParameterGroupProperties, meta)
}

func resourceNeptuneDBParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(neptuneDBParameterGroupType, data, meta)
}

func resourceNeptuneDBParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(neptuneDBParameterGroupType, data, meta)
}
