// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogLaunchTemplateConstraintType string = "AWS::ServiceCatalog::LaunchTemplateConstraint"

var serviceCatalogLaunchTemplateConstraintProperties map[string]string = map[string]string{
	"description": "Description",
	"accept_language": "AcceptLanguage",
	"portfolio_id": "PortfolioId",
	"product_id": "ProductId",
	"rules": "Rules",
}

func ResourceServiceCatalogLaunchTemplateConstraint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogLaunchTemplateConstraintExists,
		Read: resourceServiceCatalogLaunchTemplateConstraintRead,
		Create: resourceServiceCatalogLaunchTemplateConstraintCreate,
		Update: resourceServiceCatalogLaunchTemplateConstraintUpdate,
		Delete: resourceServiceCatalogLaunchTemplateConstraintDelete,
		CustomizeDiff: resourceServiceCatalogLaunchTemplateConstraintCustomizeDiff,
		
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
			},
			"product_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"rules": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogLaunchTemplateConstraintExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogLaunchTemplateConstraintType, ResourceServiceCatalogLaunchTemplateConstraint(), data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogLaunchTemplateConstraintType, ResourceServiceCatalogLaunchTemplateConstraint(), data, serviceCatalogLaunchTemplateConstraintProperties, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogLaunchTemplateConstraintType, ResourceServiceCatalogLaunchTemplateConstraint(), data, serviceCatalogLaunchTemplateConstraintProperties, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogLaunchTemplateConstraintType, data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogLaunchTemplateConstraintType, data, meta)
}
