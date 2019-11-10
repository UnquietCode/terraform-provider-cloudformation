// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-stacksetconstraint.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogStackSetConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogStackSetConstraintCreate,
		Read:   resourceServiceCatalogStackSetConstraintRead,
		Update: resourceServiceCatalogStackSetConstraintUpdate,
		Delete: resourceServiceCatalogStackSetConstraintDelete,

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
				Optional: true,
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

func resourceServiceCatalogStackSetConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::StackSetConstraint", data, meta)
}

func resourceServiceCatalogStackSetConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::StackSetConstraint", data, meta)
}

func resourceServiceCatalogStackSetConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::StackSetConstraint", data, meta)
}

func resourceServiceCatalogStackSetConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::StackSetConstraint", data, meta)
}