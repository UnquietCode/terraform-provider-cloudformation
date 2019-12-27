// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html

package servicecatalog

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const serviceCatalogPortfolioShareType string = "AWS::ServiceCatalog::PortfolioShare"

func DatasourceServiceCatalogPortfolioShare() *schema.Resource {
	return &schema.Resource{
		Read: datasourceServiceCatalogPortfolioShareRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceServiceCatalogPortfolioShareRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(serviceCatalogPortfolioShareType, DatasourceServiceCatalogPortfolioShare(), data, meta)
}
