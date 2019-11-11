// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbsubnetgroup.html

package neptune

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceNeptuneDBSubnetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNeptuneDBSubnetGroupCreate,
		Read:   resourceNeptuneDBSubnetGroupRead,
		Update: resourceNeptuneDBSubnetGroupUpdate,
		Delete: resourceNeptuneDBSubnetGroupDelete,

		Schema: map[string]*schema.Schema{
			"db_subnet_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceNeptuneDBSubnetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Neptune::DBSubnetGroup", ResourceNeptuneDBSubnetGroup(), data, meta)
}

func resourceNeptuneDBSubnetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Neptune::DBSubnetGroup", ResourceNeptuneDBSubnetGroup(), data, meta)
}

func resourceNeptuneDBSubnetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Neptune::DBSubnetGroup", ResourceNeptuneDBSubnetGroup(), data, meta)
}

func resourceNeptuneDBSubnetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Neptune::DBSubnetGroup", data, meta)
}