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

func ResourceServiceCatalogPortfolioProductAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceServiceCatalogPortfolioProductAssociationCreate,
		Read:   resourceServiceCatalogPortfolioProductAssociationRead,
		Update: resourceServiceCatalogPortfolioProductAssociationUpdate,
		Delete: resourceServiceCatalogPortfolioProductAssociationDelete,

		Schema: map[string]*schema.Schema{
			"source_portfolio_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
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
		},
	}
}

func resourceServiceCatalogPortfolioProductAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::PortfolioProductAssociation", data, meta)
}

func resourceServiceCatalogPortfolioProductAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::PortfolioProductAssociation", data, meta)
}

func resourceServiceCatalogPortfolioProductAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::PortfolioProductAssociation", data, meta)
}

func resourceServiceCatalogPortfolioProductAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::PortfolioProductAssociation", data, meta)
}