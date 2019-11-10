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

func ResourceServiceCatalogLaunchTemplateConstraint() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogLaunchTemplateConstraintCreate,
		Read:   resourceServiceCatalogLaunchTemplateConstraintRead,
		Update: resourceServiceCatalogLaunchTemplateConstraintUpdate,
		Delete: resourceServiceCatalogLaunchTemplateConstraintDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
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
			"rules": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceCatalogLaunchTemplateConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::LaunchTemplateConstraint", data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::LaunchTemplateConstraint", data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::LaunchTemplateConstraint", data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::LaunchTemplateConstraint", data, meta)
}