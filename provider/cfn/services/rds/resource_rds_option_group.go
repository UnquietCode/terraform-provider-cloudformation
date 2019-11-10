// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRDSOptionGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceRDSOptionGroupCreate,
		Read:   resourceRDSOptionGroupRead,
		Update: resourceRDSOptionGroupUpdate,
		Delete: resourceRDSOptionGroupDelete,

		Schema: map[string]*schema.Schema{
			"engine_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"major_engine_version": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"option_configurations": {
				Type: schema.TypeList,
				Elem: propertyOptionGroupOptionConfiguration(),
				Required: true,
				ForceNew: true,
			},
			"option_group_description": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceRDSOptionGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::OptionGroup", data, meta)
}

func resourceRDSOptionGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::OptionGroup", data, meta)
}

func resourceRDSOptionGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::OptionGroup", data, meta)
}

func resourceRDSOptionGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::OptionGroup", data, meta)
}