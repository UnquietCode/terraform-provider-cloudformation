// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-resourceupdateconstraint.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogResourceUpdateConstraintType string = "AWS::ServiceCatalog::ResourceUpdateConstraint"

var serviceCatalogResourceUpdateConstraintProperties map[string]string = map[string]string{
	"description": "Description",
	"accept_language": "AcceptLanguage",
	"tag_update_on_provisioned_product": "TagUpdateOnProvisionedProduct",
	"portfolio_id": "PortfolioId",
	"product_id": "ProductId",
}

func ResourceServiceCatalogResourceUpdateConstraint() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogResourceUpdateConstraintExists,
		Read: resourceServiceCatalogResourceUpdateConstraintRead,
		Create: resourceServiceCatalogResourceUpdateConstraintCreate,
		Update: resourceServiceCatalogResourceUpdateConstraintUpdate,
		Delete: resourceServiceCatalogResourceUpdateConstraintDelete,
		CustomizeDiff: resourceServiceCatalogResourceUpdateConstraintCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tag_update_on_provisioned_product": {
				Type: schema.TypeString,
				Required: true,
			},
			"portfolio_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"product_id": {
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

func resourceServiceCatalogResourceUpdateConstraintExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogResourceUpdateConstraintType, ResourceServiceCatalogResourceUpdateConstraint(), data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogResourceUpdateConstraintType, ResourceServiceCatalogResourceUpdateConstraint(), data, serviceCatalogResourceUpdateConstraintProperties, meta)
}

func resourceServiceCatalogResourceUpdateConstraintUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogResourceUpdateConstraintType, ResourceServiceCatalogResourceUpdateConstraint(), data, serviceCatalogResourceUpdateConstraintProperties, meta)
}

func resourceServiceCatalogResourceUpdateConstraintDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogResourceUpdateConstraintType, data, meta)
}

func resourceServiceCatalogResourceUpdateConstraintCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogResourceUpdateConstraintType, data, meta)
}
