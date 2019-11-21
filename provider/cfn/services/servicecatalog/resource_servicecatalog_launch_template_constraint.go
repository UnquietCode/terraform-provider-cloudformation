// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html

package servicecatalog

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogLaunchTemplateConstraintType string = "AWS::ServiceCatalog::LaunchTemplateConstraint"

func ResourceServiceCatalogLaunchTemplateConstraint() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceServiceCatalogLaunchTemplateConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogLaunchTemplateConstraintType, ResourceServiceCatalogLaunchTemplateConstraint(), data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogLaunchTemplateConstraintType, ResourceServiceCatalogLaunchTemplateConstraint(), data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogLaunchTemplateConstraintType, ResourceServiceCatalogLaunchTemplateConstraint(), data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogLaunchTemplateConstraintType, data, meta)
}

func resourceServiceCatalogLaunchTemplateConstraintCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogLaunchTemplateConstraintType, data, meta)
}
