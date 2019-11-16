// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSDBParameterGroupType string = "AWS::RDS::DBParameterGroup"

var rDSDBParameterGroupProperties map[string]string = map[string]string{
	"description": "Description",
	"family": "Family",
	"parameters": "Parameters",
	"tags": "Tags",
}

func ResourceRDSDBParameterGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRDSDBParameterGroupExists,
		Read: resourceRDSDBParameterGroupRead,
		Create: resourceRDSDBParameterGroupCreate,
		Update: resourceRDSDBParameterGroupUpdate,
		Delete: resourceRDSDBParameterGroupDelete,
		CustomizeDiff: resourceRDSDBParameterGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"family": {
				Type: schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
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

func resourceRDSDBParameterGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRDSDBParameterGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSDBParameterGroupType, ResourceRDSDBParameterGroup(), data, meta)
}

func resourceRDSDBParameterGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rDSDBParameterGroupType, ResourceRDSDBParameterGroup(), data, rDSDBParameterGroupProperties, meta)
}

func resourceRDSDBParameterGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rDSDBParameterGroupType, ResourceRDSDBParameterGroup(), data, rDSDBParameterGroupProperties, meta)
}

func resourceRDSDBParameterGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rDSDBParameterGroupType, data, meta)
}

func resourceRDSDBParameterGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rDSDBParameterGroupType, data, meta)
}
