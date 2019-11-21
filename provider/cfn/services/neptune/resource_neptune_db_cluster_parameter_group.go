// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbclusterparametergroup.html

package neptune

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const neptuneDBClusterParameterGroupType string = "AWS::Neptune::DBClusterParameterGroup"

func ResourceNeptuneDBClusterParameterGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceNeptuneDBClusterParameterGroupRead,
		Create: resourceNeptuneDBClusterParameterGroupCreate,
		Update: resourceNeptuneDBClusterParameterGroupUpdate,
		Delete: resourceNeptuneDBClusterParameterGroupDelete,
		CustomizeDiff: resourceNeptuneDBClusterParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceNeptuneDBClusterParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(neptuneDBClusterParameterGroupType, ResourceNeptuneDBClusterParameterGroup(), data, meta)
}

func resourceNeptuneDBClusterParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(neptuneDBClusterParameterGroupType, ResourceNeptuneDBClusterParameterGroup(), data, meta)
}

func resourceNeptuneDBClusterParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(neptuneDBClusterParameterGroupType, ResourceNeptuneDBClusterParameterGroup(), data, meta)
}

func resourceNeptuneDBClusterParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(neptuneDBClusterParameterGroupType, data, meta)
}

func resourceNeptuneDBClusterParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(neptuneDBClusterParameterGroupType, data, meta)
}
