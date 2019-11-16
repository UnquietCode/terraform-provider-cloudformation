// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolio.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogPortfolioType string = "AWS::ServiceCatalog::Portfolio"

var serviceCatalogPortfolioProperties map[string]string = map[string]string{
	"provider_name": "ProviderName",
	"description": "Description",
	"display_name": "DisplayName",
	"accept_language": "AcceptLanguage",
	"tags": "Tags",
}

func ResourceServiceCatalogPortfolio() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogPortfolioExists,
		Read: resourceServiceCatalogPortfolioRead,
		Create: resourceServiceCatalogPortfolioCreate,
		Update: resourceServiceCatalogPortfolioUpdate,
		Delete: resourceServiceCatalogPortfolioDelete,
		CustomizeDiff: resourceServiceCatalogPortfolioCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"provider_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogPortfolioExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogPortfolioRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogPortfolioType, ResourceServiceCatalogPortfolio(), data, meta)
}

func resourceServiceCatalogPortfolioCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogPortfolioType, ResourceServiceCatalogPortfolio(), data, serviceCatalogPortfolioProperties, meta)
}

func resourceServiceCatalogPortfolioUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogPortfolioType, ResourceServiceCatalogPortfolio(), data, serviceCatalogPortfolioProperties, meta)
}

func resourceServiceCatalogPortfolioDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogPortfolioType, data, meta)
}

func resourceServiceCatalogPortfolioCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogPortfolioType, data, meta)
}
