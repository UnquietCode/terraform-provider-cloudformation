// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogPortfolioShare() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogPortfolioShareExists,
		Read: resourceServiceCatalogPortfolioShareRead,
		Create: resourceServiceCatalogPortfolioShareCreate,
		Update: resourceServiceCatalogPortfolioShareUpdate,
		Delete: resourceServiceCatalogPortfolioShareDelete,
		CustomizeDiff: resourceServiceCatalogPortfolioShareCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"account_id": {
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
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogPortfolioShareExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogPortfolioShareRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::PortfolioShare", ResourceServiceCatalogPortfolioShare(), data, meta)
}

func resourceServiceCatalogPortfolioShareCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::PortfolioShare", ResourceServiceCatalogPortfolioShare(), data, meta)
}

func resourceServiceCatalogPortfolioShareUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::PortfolioShare", ResourceServiceCatalogPortfolioShare(), data, meta)
}

func resourceServiceCatalogPortfolioShareDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::PortfolioShare", data, meta)
}

func resourceServiceCatalogPortfolioShareCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
