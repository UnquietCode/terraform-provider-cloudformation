// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceServiceCatalogPortfolioShareCreate,
		Read:   resourceServiceCatalogPortfolioShareRead,
		Delete: resourceServiceCatalogPortfolioShareDelete,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accept_language": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"portfolio_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceServiceCatalogPortfolioShareCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::PortfolioShare", data, meta)
}

func resourceServiceCatalogPortfolioShareRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::PortfolioShare", data, meta)
}

func resourceServiceCatalogPortfolioShareUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::PortfolioShare", data, meta)
}

func resourceServiceCatalogPortfolioShareDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::PortfolioShare", data, meta)
}