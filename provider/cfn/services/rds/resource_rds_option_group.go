// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const rDSOptionGroupType string = "AWS::RDS::OptionGroup"

func ResourceRDSOptionGroup() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRDSOptionGroupExists,
		Read: resourceRDSOptionGroupRead,
		Create: resourceRDSOptionGroupCreate,
		Update: resourceRDSOptionGroupUpdate,
		Delete: resourceRDSOptionGroupDelete,
		CustomizeDiff: resourceRDSOptionGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"engine_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"major_engine_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"option_configurations": {
				Type: schema.TypeList,
				Elem: propertyOptionGroupOptionConfiguration(),
				Required: true,
			},
			"option_group_description": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRDSOptionGroupExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRDSOptionGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(rDSOptionGroupType, ResourceRDSOptionGroup(), data, meta)
}

func resourceRDSOptionGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(rDSOptionGroupType, ResourceRDSOptionGroup(), data, meta)
}

func resourceRDSOptionGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(rDSOptionGroupType, ResourceRDSOptionGroup(), data, meta)
}

func resourceRDSOptionGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(rDSOptionGroupType, data, meta)
}

func resourceRDSOptionGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(rDSOptionGroupType, data, meta)
}
