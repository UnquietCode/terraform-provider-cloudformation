// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html

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
				Optional: true,
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
			"rules": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogLaunchTemplateConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::LaunchTemplateConstraint", ResourceServiceCatalogLaunchTemplateConstraint(), data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::LaunchTemplateConstraint", ResourceServiceCatalogLaunchTemplateConstraint(), data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::LaunchTemplateConstraint", ResourceServiceCatalogLaunchTemplateConstraint(), data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::LaunchTemplateConstraint", data, meta)
}