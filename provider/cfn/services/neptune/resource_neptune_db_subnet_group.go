// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbsubnetgroup.html

package neptune

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const neptuneDBSubnetGroupType string = "AWS::Neptune::DBSubnetGroup"

func ResourceNeptuneDBSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceNeptuneDBSubnetGroupRead,
		Create: resourceNeptuneDBSubnetGroupCreate,
		Update: resourceNeptuneDBSubnetGroupUpdate,
		Delete: resourceNeptuneDBSubnetGroupDelete,
		CustomizeDiff: resourceNeptuneDBSubnetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"db_subnet_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceNeptuneDBSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(neptuneDBSubnetGroupType, ResourceNeptuneDBSubnetGroup(), data, meta)
}

func resourceNeptuneDBSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(neptuneDBSubnetGroupType, ResourceNeptuneDBSubnetGroup(), data, meta)
}

func resourceNeptuneDBSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(neptuneDBSubnetGroupType, ResourceNeptuneDBSubnetGroup(), data, meta)
}

func resourceNeptuneDBSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(neptuneDBSubnetGroupType, data, meta)
}

func resourceNeptuneDBSubnetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(neptuneDBSubnetGroupType, data, meta)
}
