// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-acceptedportfolioshare.html

package servicecatalog

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogAcceptedPortfolioShareType string = "AWS::ServiceCatalog::AcceptedPortfolioShare"

func ResourceServiceCatalogAcceptedPortfolioShare() *schema.Resource {
	return &schema.Resource{
		Read: resourceServiceCatalogAcceptedPortfolioShareRead,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceServiceCatalogAcceptedPortfolioShareRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogAcceptedPortfolioShareType, ResourceServiceCatalogAcceptedPortfolioShare(), data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(serviceCatalogAcceptedPortfolioShareType, ResourceServiceCatalogAcceptedPortfolioShare(), data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(serviceCatalogAcceptedPortfolioShareType, ResourceServiceCatalogAcceptedPortfolioShare(), data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(serviceCatalogAcceptedPortfolioShareType, data, meta)
}

func resourceServiceCatalogAcceptedPortfolioShareCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(serviceCatalogAcceptedPortfolioShareType, data, meta)
}
