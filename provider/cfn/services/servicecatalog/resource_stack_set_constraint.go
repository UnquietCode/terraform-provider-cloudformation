// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStackSetConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceStackSetConstraintCreate,
		Read:   resourceStackSetConstraintRead,
		Update: resourceStackSetConstraintUpdate,
		Delete: resourceStackSetConstraintDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: true,
			},
			"stack_instance_control": {
				Type: schema.TypeString,
				Required: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Required: false,
			},
			"portfolio_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"product_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region_list": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"admin_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"account_list": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"execution_role": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceStackSetConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::StackSetConstraint", data, meta)
}

func resourceStackSetConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::StackSetConstraint", data, meta)
}

func resourceStackSetConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::StackSetConstraint", data, meta)
}

func resourceStackSetConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::StackSetConstraint", data, meta)
}