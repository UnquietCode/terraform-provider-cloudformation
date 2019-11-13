// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceServiceCatalogAcceptedPortfolioShare() *schema.Resource {
	return &schema.Resource{
		Exists: resourceServiceCatalogAcceptedPortfolioShareExists,
		Read:   resourceServiceCatalogAcceptedPortfolioShareRead,
		Create: resourceServiceCatalogAcceptedPortfolioShareCreate,
		Update: resourceServiceCatalogAcceptedPortfolioShareUpdate,
		Delete: resourceServiceCatalogAcceptedPortfolioShareDelete,
		CustomizeDiff: resourceServiceCatalogAcceptedPortfolioShareCustomizeDiff,

		Schema: map[string]*schema.Schema{
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

func resourceServiceCatalogAcceptedPortfolioShareExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ServiceCatalog::AcceptedPortfolioShare", ResourceServiceCatalogAcceptedPortfolioShare(), data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ServiceCatalog::AcceptedPortfolioShare", ResourceServiceCatalogAcceptedPortfolioShare(), data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ServiceCatalog::AcceptedPortfolioShare", ResourceServiceCatalogAcceptedPortfolioShare(), data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ServiceCatalog::AcceptedPortfolioShare", data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}