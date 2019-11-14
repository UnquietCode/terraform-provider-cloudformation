// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogPortfolioProductAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogPortfolioProductAssociationExists,
		Read: resourceServiceCatalogPortfolioProductAssociationRead,
		Create: resourceServiceCatalogPortfolioProductAssociationCreate,
		Update: resourceServiceCatalogPortfolioProductAssociationUpdate,
		Delete: resourceServiceCatalogPortfolioProductAssociationDelete,
		CustomizeDiff: resourceServiceCatalogPortfolioProductAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"source_portfolio_id": {
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogPortfolioProductAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogPortfolioProductAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::PortfolioProductAssociation", ResourceServiceCatalogPortfolioProductAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioProductAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::PortfolioProductAssociation", ResourceServiceCatalogPortfolioProductAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioProductAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::PortfolioProductAssociation", ResourceServiceCatalogPortfolioProductAssociation(), data, meta)
}

func resourceServiceCatalogPortfolioProductAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::PortfolioProductAssociation", data, meta)
}

func resourceServiceCatalogPortfolioProductAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

